package notification_mysql

import (
	"database/sql"
	"fmt"
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

func (n *NotificationMySQL) GetByUserID(id uint) ([]notification.NotificationData, error) {
	rows, err := n.db.Query(`
		SELECT
			id,
			title,
			contents,
			created_at,
			target_path,
			read_status
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
		); err != nil {
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
			read_status
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
			read_status
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
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	notification.NotificationDB = &NotificationMySQL{
		db: db,
	}
}
