package io

import (
	"architecture/ws/services/memory"

	"github.com/gin-gonic/gin"
)

func BulkWriteInMemoryRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/memory/write", memory.GetRequestForWriteOnMemory())
	incommingRoutes.GET("/memory/:address", memory.ReadFromMemory())
}

// type BulkWriteMemoryRequest struct {
// 	Address string `json:"address"`
// 	Value   string `json:"value"`
// }

// // var dataBus *bus.DataBus
// var globalDataBus *bus.DataBus
// var globalMemory *memory.Memory
