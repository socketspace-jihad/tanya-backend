package user

import "database/sql"

var (
	UserDB = &User{}
)

type UserData struct {
	ID         uint           `json:"id" mapstructure:"id"`
	UUID       string         `json:"uuid" mapstructure:uuid`
	Email      string         `json:"email" mapstructure:"email"`
	Password   string         `json:"password,omitempty"`
	PlatformID uint8          `json:"platform_id"`
	LastLogin  string         `json:"last_login,omitempty"`
	FirstName  sql.NullString `json:"first_name"`
}

func (u *UserData) Valid() error {
	return nil
}

type UsersData []UserData

type User struct {
	Repo UserRepository
}

func (u *User) Save(user *UserData) error {
	return u.Repo.Save(user)
}

func (u *User) List() UsersData {
	return u.Repo.List()
}

func (u *User) GetByEmail(email string) (UserData, error) {
	return u.Repo.GetByEmail(email)
}

func (u *User) GetByID(id uint) (UserData, error) {
	return u.Repo.GetByID(id)
}

func init() {
	if err := userModelValidation(); err != nil {
		panic(err)
	}
}
