package mainmemory

import (
	"github.com/gin-gonic/gin"
)

func MMRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/mm", Read())
	incomingRoutes.POST("/mm", Write())
}
