package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
