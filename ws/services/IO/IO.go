package io

import (
	"architecture/ws/services/bus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BulkWriteInMemoryRoutes(incommingRoutes *gin.Engine) {
	// incommingRoutes.POST("/memory/write", Write())
}

type BulkWriteMemoryRequest struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

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
		byteData := []byte(request.Value)
		WriteOnAddress(byteAddress, byteData)
		c.JSON(http.StatusOK, gin.H{"message": "Data written successfully"})
	}
}

func WriteOnAddress(byteAddress int, data []byte) {
	dataBus := bus.NewDataBus()
	dataBus.Write(byteAddress, data)
}
