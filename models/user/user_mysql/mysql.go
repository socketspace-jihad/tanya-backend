package user_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"golang.org/x/crypto/bcrypt"
)

type UserMySQL struct {
	db *sql.DB
}

func NewUserMySQL(creds *models.DBCreds) user.UserRepository {
	return &UserMySQL{}
}

func (m *UserMySQL) ValidateConnection() error {
	return nil
}

func (m *UserMySQL) Save(u user.UserData) error {
	tx, err := m.db.BeginTx(context.Background(), nil)
	defer tx.Commit()
	if err != nil {
		return err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}
	_, err = tx.Query("INSERT INTO user (email,password,platform_id) VALUES (?,?,?)", u.Email, string(hashed), u.PlatformID)
	return err
}

func (m *UserMySQL) List() user.UsersData {
	tx, err := m.db.BeginTx(context.Background(), nil)
	defer tx.Commit()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	rows, err := tx.Query("SELECT id,email,platform_id FROM user;")
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var data user.UsersData = []user.UserData{}
	for rows.Next() {
		u := user.UserData{}
		rows.Scan(&u.ID, &u.Email, &u.PlatformID)
		data = append(data, u)
	}
	return data
}

func (m *UserMySQL) GetByID(id uint) (user.UserData, error) {
	return user.UserData{}, nil
}

func (m *UserMySQL) GetByEmail(email string) (user.UserData, error) {
	var d user.UserData
	tx, err := m.db.BeginTx(context.Background(), nil)
	defer tx.Commit()
	if err != nil {
		log.Fatalln(err)
		return d, err
	}
	rows, err := tx.Query("SELECT id,email,platform_id,password FROM user WHERE email=?", email)
	if err != nil {
		return d, err
	}
	var data user.UserData
	for rows.Next() {
		rows.Scan(&data.ID, &data.Email, &data.PlatformID, &data.Password)
	}
	return data, nil
}

func init() {
	user.RegisterUserRepo("mysql", NewUserMySQL)

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
	user.UserDB.Repo = &UserMySQL{
		db: db,
	}
}
