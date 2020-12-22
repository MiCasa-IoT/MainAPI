package controllers

import (
	"MiCasa-API/internal/firebase"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"context"
	"github.com/gin-gonic/gin"
)

// SendMessageHandler ...
// @Summary メッセージを送信する
// @Accept	 json
// @Produce  json
// @Param message body models.Message{topic} true "Topic"
// @Success 200 {object} models.FCMResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/message/send/ [post]
func SendMessageHandler(ctx *gin.Context) {
	c := context.Background()
	var params models.Message
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)

	client, err := firebase.InitFirebaseCloudMessaging(c)
	logging.PrintEror(err)
	message := firebase.CreateMessage(params)
	response, err := client.SendMulticast(c, message)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, response)
}
