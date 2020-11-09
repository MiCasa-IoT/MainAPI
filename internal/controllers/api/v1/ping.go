package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler ...
// @Summary APIの接続テスト
// @Produce  json
// @Success 200 {object} []models.Ping
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/ping/ [get]
func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
