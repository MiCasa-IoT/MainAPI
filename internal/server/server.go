package server

import (
	"MiCasa-API/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func InitServer() {
	engine := gin.New()
	InitRouter(engine)
	runServer(engine, os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
}

func runServer(engine *gin.Engine, host string, port string) {
	err := engine.Run(fmt.Sprintf("%s:%s", host, port))
	logging.PrintEror(err)
}
