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
	r.POST("/login", handlers.LoginHandler)
	r.POST("/signup", handlers.SignupHandler)

	r.POST("/postgres/signup", handlers.PosgresSignupHandler)
	r.POST("/postgres/login", handlers.PostgresLoginHandler)

	return r
}
