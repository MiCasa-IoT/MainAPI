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
	runServer(engine)
}

func runServer(engine *gin.Engine) {
	err := engine.Run(GetAddress())
	logging.PrintEror(err)
}

func GetAddress() string {
	serverAddress := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	return serverAddress
}