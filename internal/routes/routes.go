package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/dbgpt/internal/handlers"
)

func InitRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.GET("/", handlers.HomeHandler)

	return r
}
