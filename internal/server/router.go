package server

import (
	_ "MiCasa-API/docs"
	"MiCasa-API/internal/controllers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter(engine *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
				deleteGroup.DELETE("/", controllers.DeleteHandler)
			}
		}
	}
}
