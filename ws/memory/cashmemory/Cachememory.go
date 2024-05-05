package cachememory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmr CacheMemoryRequest
		var cm CashMemory
		if err := c.BindJSON(&cmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		value, ok := cm.data[cmr.address]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})

	}
}

func Write() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmr CacheMemoryRequest
		var cm CashMemory

		if err := c.BindJSON(&cmr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cm.data[cmr.address] = cmr.value

		c.String(http.StatusCreated, "Value %d written at address %bcd", cmr.value, cmr.address)
	}

}
