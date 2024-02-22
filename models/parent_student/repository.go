package parent_student

import (
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type ParentStudentRepository interface {
	GetByID(uint) (*ParentStudentData, error)
	Save(*ParentStudentData) error
	GetParentsByStudentID(uint) ([]parent_profiles.ParentProfilesData, error)
	GetStudentsByParentID(uint) ([]student_profiles.StudentProfilesData, error)
}
