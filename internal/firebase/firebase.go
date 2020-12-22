package firebase

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/api/option"
)

func InitAdminSDK(ctx context.Context) (*firebase.App, error) {
	account := option.WithCredentialsFile("configs/account.json")
	app, err := firebase.NewApp(ctx, nil, account)
	if err != nil {
		logging.PrintError(err)
		return nil, err
	}

	return app, nil
}

func InitMessaging(ctx context.Context, app *firebase.App) (*messaging.Client, error) {
	client, err := app.Messaging(ctx)
	if err != nil {
		logging.PrintError(err)
		return nil, err
	}

	return client, nil
}

func InitFirestore(ctx context.Context, app *firebase.App) (*firestore.Client, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		logging.PrintError(err)
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

func FilterNotificationTarget(ctx context.Context, client *firestore.Client, edgeId int) ([]bson.M, error){
	uuid, err := db.FilterByEdgeID(edgeId)
	logging.PrintError(err)
	for _, u := range uuid{
		println(u)
	}

	collection, err := client.CollectionGroup("users").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, c := range collection{
		println(c.Data())
	}

	return nil, nil
}
