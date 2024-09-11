package school_class_events

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/school_class"
	"github.com/socketspace-jihad/tanya-backend/models/school_rooms"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsData struct {
	ID                                    uint `json:"id"`
	*subjects.SubjectsData                `json:"subjects"`
	Name                                  string    `json:"name"`
	CreatedAt                             time.Time `json:"created_at"`
	StartDate                             time.Time `json:"start_date"`
	EndDate                               time.Time `json:"end_date"`
	school_class.SchoolClassData          `json:"school_class"`
	*teacher_profiles.TeacherProfilesData `json:"teacher_profiles"`
	school_rooms.SchoolRoomsData          `json:"school_rooms"`
	PresentTeacherProfiles                *teacher_profiles.TeacherProfilesData `json:"present_teacher_profiles"`
	PresentTeacherDate                    *time.Time                            `json:"present_teacher_date"`
	PresensiDate                          *time.Time                            `json:"presensi_date"`
}
