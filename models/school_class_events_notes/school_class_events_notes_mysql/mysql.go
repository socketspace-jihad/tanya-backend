package school_class_events_notes_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"
)

type SchoolClassEventsNotesMySQL struct {
	db *sql.DB
}

func (s *SchoolClassEventsNotesMySQL) Save(data *school_class_events_notes.SchoolClassEventsNotesData) error {
	tx, err := s.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			class_events_notes
		(
			class_events_id,
			title,
			contents,
			teacher_profiles_id	
		)
		VALUES (?,?,?,?)
	`,
		data.SchoolClassEventsData.ID,
		data.Title,
		data.Contents,
		data.TeacherProfilesData.ID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	ids, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	data.ID = uint(ids)
	tx.Commit()
	return nil
}

func (s *SchoolClassEventsNotesMySQL) GetByClassEventsID(id uint) ([]school_class_events_notes.SchoolClassEventsNotesData, error) {
	rows, err := s.db.Query(`
		SELECT
			cen.id,
			cen.title,
			cen.contents
		FROM class_events_notes AS cen
		WHERE cen.class_events_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var notes []school_class_events_notes.SchoolClassEventsNotesData
	for rows.Next() {
		var note school_class_events_notes.SchoolClassEventsNotesData
		rows.Scan(
			&note.ID,
			&note.Title,
			&note.Contents,
		)
		notes = append(notes, note)
	}
	return notes, nil
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
	school_class_events_notes.SchoolClassEventsNotesDB = &SchoolClassEventsNotesMySQL{
		db: db,
	}
}
