package events

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/socketspace-jihad/tanya-backend/notification"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
	_ "github.com/socketspace-jihad/tanya-backend/queue/engine/nats"
)

func PembuatanTugasEvents(e engine.QueueEngine, consumer Consumer) {
	data, err := e.Consume(queue.TSekolahPembuatanTugas)
	for {
		select {
		case <-err:
			log.Println(err)
			return
		case d := <-data:
			var data PembuatanTugas
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
			fmt.Println(fmt.Sprintf("%v.pembuatan-tugas-student-%v", consumer.Name, data.StudentID))
			err = notifier.Notify(&notification.Notification{
				Title: data.Message,
				Topic: fmt.Sprintf("%v.pembuatan-tugas-student-%v", consumer.Name, data.StudentID),
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
