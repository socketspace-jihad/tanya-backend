package user_roles

type UserTokenRepository interface {
	ValidateConnection() error
	Save(UserRolesData) error
	List() []UserRolesData
	GetByID(uint) (UserRolesData, error)
	GetByUserID(uint) (UserRolesData, error)
	GetByRoleID(uint) (UserRolesData, error)
	GetByRoleAndUserID(uint, uint) (UserRolesData, error)
}
