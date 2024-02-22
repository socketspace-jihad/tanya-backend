package user_roles

type UserRolesData struct {
	ID     uint `json:"id" mapstructure:"id"`
	UserID uint `json:"user_id" mapstructure:"user_id"`
	RoleID uint `json:"role_id" mapstructure:"role_id"`
}

func (u *UserRolesData) Valid() error {
	return nil
}
