package student_events_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/student_events"
)

type StudentEventMySQL struct {
	db *sql.DB
}

func (s *StudentEventMySQL) Save(data *student_events.StudentEventsData) error {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Query(`
		INSERT INTO student_events(
			subjects_id,
			name,
			start_date,
			end_date,
			student_profiles_id
		) VALUES (?,?,?,?,?)
	`,
		data.SubjectsData.ID,
		data.Name,
		data.StartDate,
		data.EndDate,
		data.StudentProfilesData.ID,
	); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *StudentEventMySQL) GetByID(id uint) (*student_events.StudentEventsData, error) {
	row, err := s.db.Query(`
		SELECT
			se.id,
			se.name,
			se.start_date,
			se.end_date
		FROM
			student_events AS se
		WHERE id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	var data student_events.StudentEventsData

	for row.Next() {
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
		); err != nil {
			return nil, err
		}
	}

	return &data, nil
}

func (s *StudentEventMySQL) GetByStudentID(id uint) ([]student_events.StudentEventsData, error) {
	row, err := s.db.Query(`
		SELECT
			se.id,
			se.name,
			se.start_date,
			se.end_date,
			subjects.name
		FROM
			student_events AS se
		LEFT JOIN subjects
			ON subjects.id = se.subjects_id
		WHERE se.student_profiles_id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	var events []student_events.StudentEventsData

	for row.Next() {
		var data student_events.StudentEventsData
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SubjectsData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
}

func (s *StudentEventMySQL) GetByStudentIDAndTimestamp(id uint, t time.Time) ([]student_events.StudentEventsData, error) {
	row, err := s.db.Query(`
	SELECT
		se.id,
		se.name,
		se.start_date,
		se.end_date,
		subjects.name
	FROM
		student_events AS se
	LEFT JOIN subjects
		ON subjects.id = se.subjects_id
	WHERE se.student_profiles_id = ? AND DATE(?) = DATE(se.end_date)
`, id, t)
	if err != nil {
		return nil, err
	}

	var events []student_events.StudentEventsData

	for row.Next() {
		var data student_events.StudentEventsData
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SubjectsData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
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
	student_events.StudentEventDB = &StudentEventMySQL{
		db: db,
	}
}
