package main

import (
	"MiCasa-API/configs"
	"MiCasa-API/internal/server"
)

func main() {
	configs.LoadConfig()
	server.InitServer()
}
