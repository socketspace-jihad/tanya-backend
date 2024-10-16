package school_class_events_notes_personal_viewer

import (
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalViewerData struct {
	ID                                                   uint `json:"id"`
	school_class_events_notes.SchoolClassEventsNotesData `json:"school_class_events_notes"`
	*teacher_profiles.TeacherProfilesData                `json:"teacher_profiles",omitempty`
	*student_profiles.StudentProfilesData                `json:"student_profiles",omitempty`
	*parent_profiles.ParentProfilesData                  `json:"parent_profiles",omitempty`
}
