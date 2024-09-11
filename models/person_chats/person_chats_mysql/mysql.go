package person_chats_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/person_chats"
)

type PersonChatsMySQL struct {
	db *sql.DB
}

func (p *PersonChatsMySQL) Save(chat *person_chats.PersonChatsData) error {
	return nil
}

func (p *PersonChatsMySQL) GetBySourceUserID(id uint) ([]person_chats.PersonChatsData, error) {
	rows, err := p.db.Query(`
		SELECT
			pc.source_user_id,
			pc.target_user_id,
			pc.global_chat_id,
			tu.first_name,
			tu.email
		FROM person_chats AS pc
		LEFT JOIN user AS tu
			ON tu.id = pc.target_user_id
		WHERE pc.source_user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	chats := []person_chats.PersonChatsData{}
	for rows.Next() {
		var chat person_chats.PersonChatsData
		if err := rows.Scan(
			&chat.SourceUser.ID,
			&chat.TargetUser.ID,
			&chat.GlobalChatsData.ID,
			&chat.TargetUser.FirstName,
			&chat.TargetUser.Email,
		); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
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
	person_chats.PersonChatsDB = &PersonChatsMySQL{
		db: db,
	}
}
