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
	logging.PrintErrorWithGinContext(err, ctx)

	app, err := firebase.InitAdminSDK(c)
	logging.PrintError(err)

	firestoreClient, err := firebase.InitFirestore(c, app)
	logging.PrintError(err)

	messagingClient, err := firebase.InitMessaging(c, app)
	logging.PrintError(err)

	tokens, err := firebase.FilterNotificationTarget(c, firestoreClient, params.EdgeID)
	logging.PrintError(err)

	if tokens != nil {
		params.Tokens = tokens
		message := firebase.CreateMessage(params)
		response, err := messagingClient.SendMulticast(c, message)

		logging.PrintErrorWithGinContext(err, ctx)
		logging.StatusOK(err, ctx, response)
	} else {
		logging.StatusBadRequest(err, ctx, "Target not found")
	}
}
