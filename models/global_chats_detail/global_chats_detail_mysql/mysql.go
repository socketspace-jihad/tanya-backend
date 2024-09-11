package global_chats_detail_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/global_chats_detail"
)

type GlobalChatsDetailMySQL struct {
	*sql.DB
}

func (g *GlobalChatsDetailMySQL) Save(data *global_chats_detail.GlobalChatsDetailData) error {
	tx, err := g.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Exec(`
		INSERT INTO
			global_chats_detail
		(
			source_user_id,
			target_user_id,
			global_chats_id,
			contents,
			attachment
		)
		VALUES
		(
			?,?,?,?,?
		)
	`,
		data.SourceUser.ID,
		data.TargetUser.ID,
		data.GlobalChatsData.ID,
		data.Contents,
		data.Attachment,
	); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (g *GlobalChatsDetailMySQL) GetByGlobalChatsID(id uint) ([]global_chats_detail.GlobalChatsDetailData, error) {
	rows, err := g.DB.Query(`
		SELECT
			source_user_id,
			contents,
			created_at
		FROM global_chats_detail
		WHERE global_chats_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	chats := []global_chats_detail.GlobalChatsDetailData{}
	for rows.Next() {
		var chat global_chats_detail.GlobalChatsDetailData
		if err := rows.Scan(
			&chat.SourceUser.ID,
			&chat.Contents,
			&chat.CreatedAt,
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
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	global_chats_detail.GlobalChatDetailsDB = &GlobalChatsDetailMySQL{
		DB: db,
	}
}
