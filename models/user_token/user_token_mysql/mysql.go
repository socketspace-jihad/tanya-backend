package user_token_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_token"
	"golang.org/x/crypto/bcrypt"
)

type UserTokenMySQL struct {
	db *sql.DB
}

func NewUserTokenMySQL(creds *models.DBCreds) user_token.UserTokenRepository {
	return &UserTokenMySQL{}
}

func (m *UserTokenMySQL) ValidateConnection() error {
	return nil
}

func (m *UserTokenMySQL) Save(u user_token.UserTokenData) error {
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

func (m *UserTokenMySQL) List() []user_token.UserTokenData {
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
	data := []user_token.UserTokenData{}
	for rows.Next() {
		u := user_token.UserTokenData{}
		rows.Scan(&u.ID, &u.Email, &u.PlatformID)
		data = append(data, u)
	}
	return data
}

func (m *UserTokenMySQL) GetByID(id uint) (user_token.UserTokenData, error) {
	return user_token.UserTokenData{}, nil
}

func (m *UserTokenMySQL) GetByToken(token string) (user_token.UserTokenData, error) {
	var d user_token.UserTokenData
	tx, err := m.db.BeginTx(context.Background(), nil)
	defer tx.Commit()
	if err != nil {
		log.Fatalln(err)
		return d, err
	}
	rows, err := tx.Query("SELECT id,token FROM user_token WHERE token=?", token)
	if err != nil {
		return d, err
	}
	var data user_token.UserTokenData
	for rows.Next() {
		rows.Scan(&data.ID, &data.Email, &data.PlatformID, &data.Password)
	}
	return data, nil
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
	user_token.UserTokenDB.UserTokenRepository = &UserTokenMySQL{
		db: db,
	}
}
