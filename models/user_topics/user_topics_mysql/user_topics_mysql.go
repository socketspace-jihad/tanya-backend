package user_topics_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/user_topics"
)

type UserTopicsMySQL struct {
	*sql.DB
}

func (u *UserTopicsMySQL) Save(data *user_topics.UserTopicsData) error {
	tx, err := u.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT IGNORE INTO
		user_topics (user_id,topic)
		VALUES (?,?)
	`, data.UserData.ID, data.Topic)
	if err != nil {
		tx.Rollback()
		return err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	data.ID = uint(ID)
	return nil
}

func (u *UserTopicsMySQL) GetByID(id uint) (*user_topics.UserTopicsData, error) {
	return nil, nil
}

func (u *UserTopicsMySQL) GetByUserID(id uint) ([]user_topics.UserTopicsData, error) {
	rows, err := u.DB.Query(`
		SELECT 
		id,
		user_id,
		topic
		FROM user_topics
		WHERE user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var data []user_topics.UserTopicsData
	for rows.Next() {
		var userTopic user_topics.UserTopicsData
		if err := rows.Scan(
			&userTopic.ID,
			&userTopic.UserData.ID,
			&userTopic.Topic,
		); err != nil {
			return nil, err
		}
		data = append(data, userTopic)
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
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	user_topics.UserTopicsDB = &UserTopicsMySQL{
		DB: db,
	}
}
