package processormemory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pmr ProcessorMemoryRequest
		var pm ProcessorMemory

		if err := c.BindJSON(&pmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		value, ok := pm.data[pmr.address]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": value})
	}
}

func Write() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pmr ProcessorMemoryRequest
		var pm ProcessorMemory

		if err := c.Bind(&pmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pm.data[pmr.address] = pmr.value

		c.String(http.StatusCreated, "Value %d written at address %bcd", pmr.value, pmr.address)

	}
}
