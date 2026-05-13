package router

import (
	"github.com/gin-gonic/gin"
	"vuegin/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)

		auth.POST("/register", controller.Register)
	}
	return r
}
