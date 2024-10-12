package event_siswa

import (
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	orangtua "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

func init() {
	log.Println("CONSUMING EVENT SISWA UNTUK ORANG TUA")
	platformFactory, err := engine.GetQueueEngine(os.Getenv("EVENT_KELAS_QUEUE_ENGINE"))
	if err != nil {
		panic(err)
	}
	platform := platformFactory()
	if err := platform.Connect(&engine.EngineAuthData{
		Host:     os.Getenv("EVENT_KELAS_QUEUE_HOST"),
		Username: os.Getenv("EVENT_KELAS_QUEUE_USERNAME"),
		Password: os.Getenv("EVENT_KELAS_QUEUE_PASSWORD"),
	}); err != nil {
		panic(err)
	}
	go events.EventSiswa(platform, orangtua.Consumer)
}
