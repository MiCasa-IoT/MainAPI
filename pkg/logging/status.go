package logging

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PrintErrorWithGinContext(err error, ctx *gin.Context) {
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
		})
	}
}

func StatusOK(err error, ctx *gin.Context, result interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

func StatusBadRequest(err error, ctx *gin.Context, result interface{}) {
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
		})
	}
}