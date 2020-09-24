package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
