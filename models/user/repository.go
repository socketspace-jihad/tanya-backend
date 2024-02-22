package user

import (
	"errors"

	"github.com/socketspace-jihad/tanya-backend/models"
)

type UserRepository interface {
	ValidateConnection() error
	Save(UserData) error
	List() UsersData
	GetByID(uint) (UserData, error)
	GetByEmail(string) (UserData, error)
}

type UserRepoFactory func(creds *models.DBCreds) UserRepository

var userRepoMap map[string]UserRepoFactory = make(map[string]UserRepoFactory)

func RegisterUserRepo(name string, fact UserRepoFactory) {
	userRepoMap[name] = fact
}

func GetUserRepo(name string, creds *models.DBCreds) (UserRepoFactory, error) {
	if _, ok := userRepoMap[name]; !ok {
		return nil, errors.New("user repo not found, couldn't find any implementation for " + name + " db")
	}
	if err := userRepoMap[name](creds).ValidateConnection(); err != nil {
		return nil, err
	}
	return userRepoMap[name], nil
}
