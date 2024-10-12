package siswa

import (
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
)

var (
	Consumer events.Consumer = events.Consumer{
		Name: "siswa",
	}
)
