package pembuatan_tugas

import (
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	orangtua "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

func init() {
	log.Println("CONSUMING PEMBUATAN TUGAS UNTUK ORANG TUA")
	platformFactory, err := engine.GetQueueEngine(os.Getenv("PEMBUATAN_TUGAS_QUEUE_ENGINE"))
	if err != nil {
		panic(err)
	}
	platform := platformFactory()
	if err := platform.Connect(&engine.EngineAuthData{
		Host:     os.Getenv("PEMBUATAN_TUGAS_QUEUE_HOST"),
		Username: os.Getenv("PEMBUATAN_TUGAS_QUEUE_USERNAME"),
		Password: os.Getenv("PEMBUATAN_TUGAS_QUEUE_PASSWORD"),
	}); err != nil {
		panic(err)
	}
	go events.PembuatanTugasEvents(platform, orangtua.Consumer)
}
