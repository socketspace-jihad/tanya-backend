package presensi_siswa

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/notification"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa"
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
				Title: "Putra/i Bapak/Ibu baru saja melakukan presensi Kelas",
				Topic: fmt.Sprintf("%v.presensi-%v", siswa.Consumer.Name, data.StudentProfilesData.ID),
				Data: map[string]string{
					"created_at": data.CreatedAt.String(),
					"name":       data.StudentProfilesData.Name,
				},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func init() {
	log.Println("CONSUMING PRESENSI KELAS UNTUK SISWA")
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
