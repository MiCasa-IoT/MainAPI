package server

import (
	"MiCasa-API/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		hello := v1.Group("/ping")
		{
			hello.GET("/", controllers.PingHandler)
		}
		db := v1.Group("/db")
		{
			document := db.Group("/document")
			{
				findAllGroup := document.Group("/readall")
				findAllGroup.GET("/", controllers.ReadAllDocumentHandler)
				createGroup := document.Group("/create")
				createGroup.POST("/", controllers.CreateHandler)
				readGroup := document.Group("/read")
				readGroup.POST("/", controllers.ReadHandler)
				updateGroup := document.Group("/update")
				updateGroup.POST("/", controllers.UpdateHandler)
				deleteGroup := document.Group("/delete")
				deleteGroup.POST("/", controllers.DeleteHandler)
			}
		}
	}
}
