package events

import (
	"encoding/json"
	"fmt"
	"log"

	notification_model "github.com/socketspace-jihad/tanya-backend/models/notification"
	"github.com/socketspace-jihad/tanya-backend/notification"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
	_ "github.com/socketspace-jihad/tanya-backend/queue/engine/nats"
)

type EventKelasData struct {
	ClassID  uint
	Title    string
	Subtitle string
	notification_model.NotificationData
}

func EventKelas(e engine.QueueEngine, consumer Consumer) {
	data, err := e.Consume(queue.TEventKelas)
	for {
		select {
		case <-err:
			log.Println(err)
			return
		case d := <-data:
			var data EventKelasData
			dataBytes, _ := d.([]byte)
			if err := json.Unmarshal(dataBytes, &data); err != nil {
				log.Println("ERROR UNMARSHAL", err)
				return
			}
			notifierFactory, err := notification.GetNotifier("fcm")
			if err != nil {
				log.Println(err)
				return
			}
			notifier := notifierFactory()
			err = notifier.Notify(&notification.Notification{
				Title:    data.Title,
				Subtitle: data.Subtitle,
				Topic:    fmt.Sprintf("%v.event-kelas-%v", consumer.Name, data.ClassID),
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
