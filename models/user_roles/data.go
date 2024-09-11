package user_roles

import (
	"github.com/socketspace-jihad/tanya-backend/models/roles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type UserRolesData struct {
	ID                       uint `json:"id" mapstructure:"id"`
	user.UserData            `json:"user" mapstructure:"user"`
	roles.RolesData          `json:"roles" mapstructure:"roles"`
	CurrentStudentProfilesID uint `json:"current_student_profiles_id"`
	CurrentTeacherProfilesID uint `json:"current_teacher_profiles_id"`
}

func (u *UserRolesData) Valid() error {
	return nil
}
