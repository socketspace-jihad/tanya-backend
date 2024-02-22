package teacher_profiles

import (
	"github.com/socketspace-jihad/tanya-backend/models/schools"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type TeacherProfilesData struct {
	ID                       uint   `json:"id"`
	Name                     string `json:"name"`
	user_roles.UserRolesData `json:"user_roles"`
	schools.SchoolData       `json:"school"`
}
