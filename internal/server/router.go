package server

import (
	"MiCasa-API/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		hello := v1.Group("/hello")
		{
			hello.GET("/", controllers.HelloHandler)
		}
		db := v1.Group("/db")
		{
			document := db.Group("/document")
			{
				findAll := document.Group("/findall")
				findAll.GET("/", controllers.FindAllDocumentHandler)
			}

		}
	}
}
