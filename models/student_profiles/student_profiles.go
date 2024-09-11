package student_profiles

import (
	"github.com/socketspace-jihad/tanya-backend/models/school_class"
	"github.com/socketspace-jihad/tanya-backend/models/schools"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
	"github.com/socketspace-jihad/tanya-backend/models/verified_status"
)

type StudentProfilesData struct {
	ID                       uint `json:"id"`
	user_roles.UserRolesData `json:"user_roles" mapstructure:"user_roles"`
	NISN                     string `json:"nisn"`
	schools.SchoolData       `json:"school"`
	Name                     string                         `json:"name"`
	CurrentSchoolData        school_class.SchoolClassData   `json:"current_school_class"`
	FirstName                string                         `json:"first_name"`
	LastName                 string                         `json:"last_name"`
	Batch                    uint                           `json:"batch"`
	VerifiedStatus           verified_status.VerifiedStatus `json:"verified_status"`
}
