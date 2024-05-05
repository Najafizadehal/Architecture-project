package cachememory

import "github.com/gin-gonic/gin"

func CasheRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/cm")
}
