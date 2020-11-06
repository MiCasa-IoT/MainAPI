package main

import (
	"MiCasa-API/configs"
	_ "MiCasa-API/docs"
	"MiCasa-API/internal/server"
)

// @title MiCasa Main API
// @version 1.0
// @description MiCasaのAPIドキュメント

// @contact.name Sw-Saturn

// @license.name MIT License
// @license.url https://github.com/MiCasa-IoT/MainAPI/blob/master/LICENSE

// @host localhost:8080
// @BasePath /
func main() {
	configs.LoadConfig()
	server.InitServer()
}
