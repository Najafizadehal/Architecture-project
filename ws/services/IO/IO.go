package io

import (
	"architecture/ws/services/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BulkWriteInMemoryRoutes(incommingRoutes *gin.Engine) {
	// incommingRoutes.POST("/memory/write", Write())
}

type BulkWriteMemoryRequest struct {
	address string
	value   string
}

func WriteOnBus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request BulkWriteMemoryRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		byteAddress := []byte(request.address)
		WriteOnAddress(byteAddress)
		byteData := []byte(request.value)
		WriteValueOnBus(byteData)
	}
}
func WriteOnAddress(byteAddress []byte) {
	address := bus.NewAddressBus(len(byteAddress))
	for _, b := range byteAddress {
		address.Write(int(b))
	}
}

func WriteValueOnBus(value []byte) {
	data := bus.NewDataBus(len(value))
	for _, i := range value {
		data.Write(i)
	}
}
