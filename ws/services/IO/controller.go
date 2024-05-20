package controllers

import (
	"architecture/ws/services/controlunit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ControlUnit *controlunit.ControlUnit
}

func NewController(cu *controlunit.ControlUnit) *Controller {
	return &Controller{ControlUnit: cu}
}

func (ctr *Controller) Fetch(c *gin.Context) {
	err := ctr.ControlUnit.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fetch stage completed", "IR": ctr.ControlUnit.Registers.IR})
}

func (ctr *Controller) Decode(c *gin.Context) {
	opcode, address, err := ctr.ControlUnit.Decode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Decode stage completed", "Opcode": opcode, "Address": address})
}

func (ctr *Controller) Execute(c *gin.Context) {
	opcode := c.Query("opcode")
	addressStr := c.Query("address")
	address, err := strconv.Atoi(addressStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
		return
	}
	err = ctr.ControlUnit.Execute(opcode, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Execute stage completed", "Result": ctr.ControlUnit.Registers.AC})
}

func (ctr *Controller) RunCycle(c *gin.Context) {
	err := ctr.ControlUnit.RunCycle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Cycle completed",
		"IR":        ctr.ControlUnit.Registers.IR,
		"Opcode":    ctr.ControlUnit.Registers.IR >> 12,
		"Address":   ctr.ControlUnit.Registers.IR & 0x0FFF,
		"Result":    ctr.ControlUnit.Registers.AC,
		"Registers": ctr.ControlUnit.Registers,
	})
}

func (ctr *Controller) LoadInstruction(c *gin.Context) {
	var instruction struct {
		Address string `json:"address"`
		Value   string `json:"value"`
	}
	if err := c.BindJSON(&instruction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := strconv.ParseInt(instruction.Address, 16, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
		return
	}
	value, err := strconv.ParseInt(instruction.Value, 16, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
		return
	}

	if err := ctr.ControlUnit.Memory.Write(int(address), int(value)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Instruction loaded successfully"})
}

func (ctr *Controller) LoadInstructions(c *gin.Context) {
	var instructions []struct {
		Address int `json:"address"`
		Value   int `json:"value"`
	}
	if err := c.BindJSON(&instructions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, instruction := range instructions {
		if err := ctr.ControlUnit.Memory.Write(instruction.Address, instruction.Value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Instructions loaded successfully"})
}

func (ctr *Controller) WriteRegister(c *gin.Context) {
	var req struct {
		Register string `json:"register"`
		Value    int    `json:"value"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Register {
	case "AR":
		ctr.ControlUnit.Registers.AR = req.Value
	case "PC":
		ctr.ControlUnit.Registers.PC = req.Value
	case "DR":
		ctr.ControlUnit.Registers.DR = req.Value
	case "AC":
		ctr.ControlUnit.Registers.AC = req.Value
	case "IR":
		ctr.ControlUnit.Registers.IR = req.Value
	case "TR":
		ctr.ControlUnit.Registers.TR = req.Value
	case "INPR":
		ctr.ControlUnit.Registers.INPR = req.Value
	case "OUTR":
		ctr.ControlUnit.Registers.OUTR = req.Value
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register updated successfully"})
}
func (ctr *Controller) ReadMemory(c *gin.Context) {
	addressStr := c.Param("address")
	address, err := strconv.Atoi(addressStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
		return
	}

	value, err := ctr.ControlUnit.Memory.Read(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address, "value": value})
}

func (ctr *Controller) ReadRegister(c *gin.Context) {
	register := c.Query("register")

	var value int
	switch register {
	case "AR":
		value = ctr.ControlUnit.Registers.AR
	case "PC":
		value = ctr.ControlUnit.Registers.PC
	case "DR":
		value = ctr.ControlUnit.Registers.DR
	case "AC":
		value = ctr.ControlUnit.Registers.AC
	case "IR":
		value = ctr.ControlUnit.Registers.IR
	case "TR":
		value = ctr.ControlUnit.Registers.TR
	case "INPR":
		value = ctr.ControlUnit.Registers.INPR
	case "OUTR":
		value = ctr.ControlUnit.Registers.OUTR
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"register": register, "value": value})
}
