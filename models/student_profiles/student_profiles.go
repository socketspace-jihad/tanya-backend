package student_profiles

import (
	"github.com/socketspace-jihad/tanya-backend/models/school_class"
	"github.com/socketspace-jihad/tanya-backend/models/schools"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type StudentProfilesData struct {
	ID                       uint `json:"id"`
	user_roles.UserRolesData `json:"user_roles"`
	UserID                   uint   `json:"user_id"`
	NISN                     string `json:"nisn"`
	schools.SchoolData       `json:"school"`
	Name                     string                       `json:"name"`
	CurrentSchoolData        school_class.SchoolClassData `json:"current_school_class"`
	FirstName                string                       `json:"first_name"`
	LastName                 string                       `json:"last_name"`
}
