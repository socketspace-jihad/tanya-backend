package student_events

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
)

type StudentEventsData struct {
	ID                                   uint `json:"student_events"`
	subjects.SubjectsData                `json:"subjects"`
	Name                                 string    `json:"name"`
	CreatedAt                            time.Time `json:"created_at"`
	StartDate                            time.Time `json:"start_date"`
	EndDate                              time.Time `json:"end_date"`
	student_profiles.StudentProfilesData `json:"student_profiles"`
}
