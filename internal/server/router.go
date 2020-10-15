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
				findAll := document.Group("/readall")
				findAll.GET("/", controllers.ReadAllDocumentHandler)
				create := document.Group("/create")
				create.POST("/", controllers.CreateHandler)
				read := document.Group("/read")
				read.POST("/", controllers.ReadHandler)
				update := document.Group("/update")
				update.POST("/", controllers.UpdateHandler)
				delete := document.Group("/delete")
				delete.POST("/", controllers.DeleteHandler)
			}
		}
	}
}
