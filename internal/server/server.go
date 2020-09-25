package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func InitServer() {
	engine := gin.Default()
	InitRouter(engine)
	runServer(engine, os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
}

func runServer(engine *gin.Engine, host string, port string) {
	err := engine.Run(fmt.Sprintf("%s:%s",host, port))
	if err != nil {
		fmt.Println("Gin Error:", err)
	}
}
