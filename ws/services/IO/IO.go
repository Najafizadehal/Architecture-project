package io

import (
	"architecture/ws/services/bus"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BulkWriteInMemoryRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/memory/write", getRequestForWriteOnMemory())
	incommingRoutes.GET("/memory/:address", ReadFromBus())
}

type BulkWriteMemoryRequest struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

// var dataBus *bus.DataBus
var globalDataBus *bus.DataBus

func getRequestForWriteOnMemory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request BulkWriteMemoryRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		byteAddress, err := strconv.Atoi(request.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
			return
		}
		byteData, err := hex.DecodeString(request.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
			return
		}

		if globalDataBus == nil {
			globalDataBus = bus.NewDataBus()
		}
		// ذخیره داده در DataBus
		globalDataBus.Write(byteAddress, byteData)

		// WriteOnAddress(byteAddress, byteData)
		// WriteOnAddress(byteAddress, byteData)
		c.JSON(http.StatusOK, gin.H{"message": "Data written successfully"})
	}
}

func WriteOnAddress(byteAddress int, data []byte) *bus.DataBus {
	dataBus := bus.NewDataBus()
	dataBus.Write(byteAddress, data)

	return dataBus
}

func ReadFromBus() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")

		addr, err := strconv.Atoi(address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
			return
		}

		data := globalDataBus.Read(addr)

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// bus := bus.NewDataBus()
// 		retrievedData := bus.Read(byteAddress)
// 		memory := memory.NewMemory()
// 		memory.Store(byteAddress, retrievedData)
