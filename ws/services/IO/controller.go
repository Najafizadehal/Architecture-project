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
		Address int `json:"address"`
		Value   int `json:"value"`
	}
	if err := c.BindJSON(&instruction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctr.ControlUnit.Memory.Write(instruction.Address, instruction.Value); err != nil {
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
