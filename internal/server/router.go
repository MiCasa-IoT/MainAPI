package server

import (
	"MiCasa-API/docs"
	_ "MiCasa-API/docs"
	"MiCasa-API/internal/controllers/api/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
)

func InitRouter(engine *gin.Engine) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	url := ginSwagger.URL(
		fmt.Sprintf("%s/swagger/doc.json",
			GetAddress(),
		))
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
		message := v1.Group("/message")
		{
			sendGroup := message.Group("/send")
			sendGroup.POST("/", controllers.SendMessageHandler)
		}
	}
}
