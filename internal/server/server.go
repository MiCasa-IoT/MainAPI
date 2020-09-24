package server

import "github.com/gin-gonic/gin"

func InitServer() {
	engine := gin.Default()
	InitRouter(engine)
	engine.Run(":8080")
}
