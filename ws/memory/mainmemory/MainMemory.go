package mainmemory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		var mm MainMemory
		var mmr MainMemoryRequest
		if err := c.BindJSON(&mmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		value, ok := mm.data[mmr.address]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	}
}

func Write() gin.HandlerFunc {
	return func(c *gin.Context) {
		var mm MainMemory
		var mmr MainMemoryRequest
		if err := c.BindJSON(&mmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mm.data[mmr.address] = mmr.value

		c.String(http.StatusOK, "Value %d written at address %bcd", mmr.value, mmr.address)
	}
}
