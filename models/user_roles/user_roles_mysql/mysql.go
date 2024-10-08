package user_roles_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type UserRolesMySQL struct {
	db *sql.DB
}

func (u *UserRolesMySQL) ValidateConnection() error {
	return nil
}
func (u *UserRolesMySQL) Save(userRoles user_roles.UserRolesData) error {
	return nil
}
func (u *UserRolesMySQL) List() []user_roles.UserRolesData {
	return nil
}
func (u *UserRolesMySQL) GetByID(id uint) (user_roles.UserRolesData, error) {
	return user_roles.UserRolesData{}, nil
}

func (u *UserRolesMySQL) GetByUserID(userId uint) (user_roles.UserRolesData, error) {
	return user_roles.UserRolesData{}, nil
}

func (u *UserRolesMySQL) GetByRoleID(userId uint) (user_roles.UserRolesData, error) {
	return user_roles.UserRolesData{}, nil
}

func (u *UserRolesMySQL) GetByRoleAndUserID(roleId, userId uint) (user_roles.UserRolesData, error) {
	var userRoleData user_roles.UserRolesData
	tx, err := u.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		return userRoleData, err
	}
	rows, err := tx.Query("SELECT id,user_id,roles_id FROM user_roles WHERE user_id=? AND roles_id=?", userId, roleId)
	if err != nil {
		return userRoleData, err
	}
	for rows.Next() {
		rows.Scan(&userRoleData.ID, &userRoleData.UserID, &userRoleData.RoleID)
	}
	if err := tx.Commit(); err != nil {
		return userRoleData, err
	}
	return userRoleData, nil
}

func (u *UserRolesMySQL) Valid() error {
	return nil
}

func init() {
	creds := &models.DBCreds{
		Username: os.Getenv("DATABASE_USERNAME"),
		Host:     os.Getenv("DATABASE_HOST"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Database: os.Getenv("DATABASE_DATABASE"),
		Port:     os.Getenv("DATABASE_PORT"),
	}
	_, err := user.GetUserRepo(os.Getenv("DATABASE_ENGINE"), creds)
	if err != nil {
		// panic(err)
		log.Println(err)
	}
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	user_roles.UserRolesDB = &UserRolesMySQL{
		db: db,
	}
}
