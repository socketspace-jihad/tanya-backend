package school_class_events_notes_personal_comments

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalCommentsData struct {
	ID                                                   uint `json:"id"`
	*teacher_profiles.TeacherProfilesData                `json:"teacher_profiles"`
	*parent_profiles.ParentProfilesData                  `json:"parent_profiles"`
	*student_profiles.StudentProfilesData                `json:"student_profiles"`
	Content                                              string    `json:"content"`
	CreatedAt                                            time.Time `json:"created_at"`
	school_class_events_notes.SchoolClassEventsNotesData `json:"school_class_events_notes"`
}
