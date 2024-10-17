package events

import (
	"log"
	"sync"

	"github.com/socketspace-jihad/tanya-backend/models/notification"
	"github.com/socketspace-jihad/tanya-backend/models/user_topics"
)

func createNotification(topic string, data *notification.NotificationData) {
	users, err := user_topics.UserTopicsDB.GetByTopicName(topic)
	if err != nil {
		return
	}
	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go func(user user_topics.UserTopicsData, data *notification.NotificationData) {
			defer wg.Done()
			d := *data
			d.UserData = user.UserData
			if err := notification.NotificationDB.Save(&d); err != nil {
				log.Println(err)
			}
		}(user, data)
	}
	wg.Wait()
}
