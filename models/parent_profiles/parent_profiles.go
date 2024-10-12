package parent_profiles

import (
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type ParentProfilesData struct {
	ID                       uint   `json:"id",omitempty`
	Name                     string `json:"name"`
	NIK                      string `json:"nik"`
	user_roles.UserRolesData `json:"user_roles"`
}
