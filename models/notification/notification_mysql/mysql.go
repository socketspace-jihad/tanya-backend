package notification_mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
)

type NotificationMySQL struct {
	db *sql.DB
}

func (n *NotificationMySQL) GetByID(id uint) (*notification.NotificationData, error) {
	return nil, nil
}

func (n *NotificationMySQL) Save(data *notification.NotificationData) error {
	tx, err := n.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			notification
		(
			title,
			contents,
			target_path,
			user_id,
			data
		)
		VALUES (?,?,?,?,?)
	`,
		data.Title,
		data.Contents,
		data.TargetPath,
		data.UserData.ID,
		data.Data,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	data.ID = uint(id)
	tx.Commit()
	return nil
}

func (n *NotificationMySQL) UpdateRead(data *notification.NotificationData) error {
	tx, err := n.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Exec(`
		UPDATE notification
		SET read_status = 1
		WHERE id = ? AND user_id = ?
	`, data.ID, data.UserData.ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (n *NotificationMySQL) GetByUserID(id uint) ([]notification.NotificationData, error) {
	rows, err := n.db.Query(`
		SELECT
			id,
			title,
			contents,
			created_at,
			target_path,
			read_status,
			data
		FROM notification
		WHERE user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var notifications []notification.NotificationData
	for rows.Next() {
		var notification notification.NotificationData
		if err := rows.Scan(
			&notification.ID,
			&notification.Title,
			&notification.Contents,
			&notification.CreatedAt,
			&notification.TargetPath,
			&notification.ReadStatus,
			&notification.Data,
		); err != nil {
			log.Println("ERR", err)
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (n *NotificationMySQL) GetByStudentProfilesID(id uint) ([]notification.NotificationData, error) {
	rows, err := n.db.Query(`
		SELECT
			id,
			title,
			contents,
			created_at,
			target_path,
			read_status,
			data
		FROM notification
		WHERE student_profiles_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var notifications []notification.NotificationData
	for rows.Next() {
		var notification notification.NotificationData
		if err := rows.Scan(
			&notification.ID,
			&notification.Title,
			&notification.Contents,
			&notification.CreatedAt,
			&notification.TargetPath,
			&notification.ReadStatus,
			&notification.Data,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (n *NotificationMySQL) GetByUserOrStudentProfilesID(studentId uint, userId uint) ([]notification.NotificationData, error) {
	rows, err := n.db.Query(`
		SELECT
			id,
			title,
			contents,
			target_path,
			read_status,
			data
		FROM notification
		WHERE student_profiles_id = ? OR user_id = ?
	`, studentId, userId)
	if err != nil {
		return nil, err
	}
	var notifications []notification.NotificationData
	for rows.Next() {
		var notification notification.NotificationData
		if err := rows.Scan(
			&notification.ID,
			&notification.Title,
			&notification.Contents,
			&notification.TargetPath,
			&notification.ReadStatus,
			&notification.Data,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
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
	notification.NotificationDB = &NotificationMySQL{
		db: db,
	}
}
