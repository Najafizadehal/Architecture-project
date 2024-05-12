package io

import (
	"architecture/ws/services/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BulkWriteInMemoryRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/memory/write", Write())
}

type BulkWriteMemoryRequest struct {
	address string
	value   string
}

func Write() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request BulkWriteMemoryRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		byteAddress := []byte(request.address)
		address := bus.NewAddressBus(len(byteAddress))

		for _, b := range byteAddress {
			address.Write(int(b))
		}
	}
}
