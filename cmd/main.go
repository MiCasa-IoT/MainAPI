package main

import (
	"MiCasa-API/configs"
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/server"
)

func main() {
	configs.LoadConfig()
	db.Connect()
	server.InitServer()
}
