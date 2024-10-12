package school_class_events_notes_personal

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalData struct {
	ID                                        uint `json:"id"`
	school_class_events.SchoolClassEventsData `json:"school_class_events"`
	teacher_profiles.TeacherProfilesData      `json:"teacher_profiles"`
	student_profiles.StudentProfilesData      `json:"student_profiles"`
	CreatedAt                                 time.Time `json:"created_at"`
	Judul                                     string    `json:"judul"`
	Deskripsi                                 string    `json:"deskripsi"`
}
