package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	return Router
}
