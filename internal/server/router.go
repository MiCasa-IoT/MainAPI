package server

import (
	"MiCasa-API/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		create := v1.Group("/hello")
		{
			create.GET("/", controllers.HelloHandler)
		}
	}
}