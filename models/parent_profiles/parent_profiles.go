package parent_profiles

import (
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type ParentProfilesData struct {
	ID                       uint   `json:"id"`
	Name                     string `json:"name"`
	user_roles.UserRolesData `json:"user_roles"`
}
