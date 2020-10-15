package logging

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PrintEror(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PrintErorWithGinContext(err error, ctx *gin.Context) {
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": err,
		})
	}
}

func FatalError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func StatusOK(err error, ctx *gin.Context, result interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}
