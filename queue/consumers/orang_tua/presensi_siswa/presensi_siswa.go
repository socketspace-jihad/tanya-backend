package presensi_siswa

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/notification"
	"github.com/socketspace-jihad/tanya-backend/queue"
	orangtua "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

func consumePresensiSiswa(e engine.QueueEngine) {
	data, err := e.Consume(queue.TSekolahPresensiSiswa)
	for {
		select {
		case <-err:
			log.Println(err)
			return
		case d := <-data:
			var data student_presensi.StudentPresensiData
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
				Title:    fmt.Sprintf("%v baru saja melakukan presensi Kelas", data.StudentProfilesData.Name),
				Subtitle: "Ketuk untuk melihat di mana lokasinya",
				Topic:    fmt.Sprintf("%v.presensi-student-%v", orangtua.Consumer.Name, data.StudentProfilesData.ID),
				Data: map[string]string{
					"created_at": data.CreatedAt.String(),
					"name":       data.StudentProfilesData.Name,
					"page":       "notification",
					"title":      fmt.Sprintf("%v baru saja melakukan presensi Kelas", data.StudentProfilesData.Name),
					"subtitle":   "Ketuk untuk melihat di mana lokasinya",
				},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func init() {
	log.Println("CONSUMING PRESENSI KELAS UNTUK ORANG TUA")
	platformFactory, err := engine.GetQueueEngine(os.Getenv("PRESENSI_SISWA_QUEUE_ENGINE"))
	if err != nil {
		panic(err)
	}
	platform := platformFactory()
	if err := platform.Connect(&engine.EngineAuthData{
		Host:     os.Getenv("PRESENSI_SISWA_QUEUE_HOST"),
		Username: os.Getenv("PRESENSI_SISWA_QUEUE_USERNAME"),
		Password: os.Getenv("PRESENSI_SISWA_QUEUE_PASSWORD"),
	}); err != nil {
		panic(err)
	}
	go consumePresensiSiswa(platform)
}
