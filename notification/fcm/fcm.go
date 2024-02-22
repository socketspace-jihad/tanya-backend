package fcm

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/socketspace-jihad/tanya-backend/notification"
	"google.golang.org/api/option"
)

type FCMNotification struct {
	*firebase.App
}

func NewFCMNotification() notification.Notifier {
	opt := option.WithCredentialsFile("tanya-platform-firebase-adminsdk-wvcba-6d0a54f5e3.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	return &FCMNotification{
		App: app,
	}
}

func (n *FCMNotification) Notify(msg *notification.Notification) error {
	client, err := n.App.Messaging(context.Background())
	if err != nil {
		return err
	}
	res, err := client.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Subtitle,
		},
		Data: map[string]string{
			"title":    msg.Title,
			"subtitle": msg.Subtitle,
		},
		Topic: msg.Topic,
	})
	log.Println(res)
	return err
}

func init() {
	notification.RegisterNotifier("fcm", NewFCMNotification)
}
