package firebase

import (
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func InitFirebaseCloudMessaging(ctx context.Context) (*messaging.Client, error) {
	account := option.WithCredentialsFile("configs/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, account)
	if err != nil {
		logging.PrintEror(err)
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		logging.PrintEror(err)
		return nil, err
	}

	return client, nil
}

func CreateMessage(params models.Message) *messaging.MulticastMessage {
	android := new(messaging.AndroidConfig)
	android.Priority = "high"
	androidNotification := new(messaging.AndroidNotification)
	androidNotification.ChannelID = "channel_1"
	androidNotification.Tag = params.Tag
	notification := new(messaging.Notification)
	notification.Title = params.Title
	notification.Body = params.Body

	message := &messaging.MulticastMessage {
		Data: map[string] string {
			"title": params.Title,
			"body": params.Body,
		},
		Android: android,
		Notification: notification,
		Tokens: params.Tokens,
	}
	return message
}
