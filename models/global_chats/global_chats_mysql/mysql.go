package global_chats_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/global_chats"
)

type GlobalChatsMySQL struct {
	db *sql.DB
}

func (g *GlobalChatsMySQL) Save(chat *global_chats.GlobalChatsData) error {
	return nil
}

func (g *GlobalChatsMySQL) GetBySourceUserID(id uint) ([]global_chats.GlobalChatsData, error) {
	rows, err := g.db.Query(`
		SELECT
			pc.source_user_id,
			pc.target_user_id,
			pc.global_chat_id,
		FROM person_chats AS pc
		WHERE pc.source_user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	chats := []global_chats.GlobalChatsData{}
	for rows.Next() {
		if err := rows.Scan(); err != nil {
			return nil, err
		}
	}

	return chats, nil
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
	global_chats.GlobalChatsDB = &GlobalChatsMySQL{
		db: db,
	}
}
