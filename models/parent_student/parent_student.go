package parent_student

import (
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type ParentStudentData struct {
	ID                                   uint `json:"parent_student"`
	parent_profiles.ParentProfilesData   `json:"parent_profiles"`
	student_profiles.StudentProfilesData `json:"student_profiles"`
}
