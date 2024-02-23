package student_presensi

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/event_types"
	"github.com/socketspace-jihad/tanya-backend/models/presensi_types"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type StudentPresensiData struct {
	ID                                         uint `json:"id"`
	*student_profiles.StudentProfilesData      `json:"student_profiles"`
	EventsID                                   uint `json:"events_id"`
	event_types.EventTypesData                 `json:"event_types"`
	CreatedAt                                  time.Time `json:"created_at"`
	Lattitude                                  *float32  `json:"lattitude"`
	Longitude                                  *float32  `json:"longitude"`
	*school_class_events.SchoolClassEventsData `json:"school_class_events"`
	*presensi_types.PresensitypesData          `json:"presensi_types"`
}
