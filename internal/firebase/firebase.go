package firebase

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

func InitAdminSDK(ctx context.Context) (*firebase.App, error) {
	credDec, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_APPLICATION_CREDENTIAL"))
	opt := option.WithCredentialsJSON(credDec)
	app, err := firebase.NewApp(ctx, nil, opt)
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
		logging.PrintError(fmt.Errorf("firestore: %s", err))
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
	notification.Body = fmt.Sprintf("Edge ID: %d\n%s",params.EdgeID, params.Body)

	message := &messaging.MulticastMessage {
		Data: map[string] string {
			"title": params.Title,
			"body": fmt.Sprintf("Edge ID: %d\n%s",params.EdgeID, params.Body),
		},
		Android: android,
		Notification: notification,
		Tokens: params.Tokens,
	}

	return message
}

func FilterNotificationTarget(ctx context.Context, client *firestore.Client, edgeId int) ([]string, error){
	uuid, err := db.FilterByEdgeID(edgeId)
	logging.PrintError(err)

	uuid, err = db.GetLatestUUIDToFilterByEdgeID(uuid, edgeId)
	logging.PrintError(err)

	if len(uuid) > 0 {
		collection, err := client.Collection("users").Where("uuid", "in", uuid).Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}

		var tokens []string
		for _, snapshot := range collection {
			tokens = append(tokens, snapshot.Data()["token"].(string))
		}
		return tokens, nil
	}
	return nil, nil
}
