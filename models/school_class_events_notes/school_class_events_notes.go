package school_class_events_notes

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/school_class"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesData struct {
	ID                                   uint `json:"id"`
	school_class.SchoolClassData         `json:"school_class_events"`
	Title                                string `json:"title"`
	Contents                             string `json:"contents"`
	teacher_profiles.TeacherProfilesData `json:"teacher_profiles"`
	CreatedAt                            time.Time `json:"created_at"`
	Photo                                string    `json:"photo"`
}
