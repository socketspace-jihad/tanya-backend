package school_class_events_notes_personal_pictures_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_pictures"
)

type SchoolClassEventsNotesPersonalPicturesMySQL struct {
	db *sql.DB
}

func (s *SchoolClassEventsNotesPersonalPicturesMySQL) Save(data *school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesData) error {
	tx, err := s.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			class_events_notes_personal_pictures
		(
			path,
			class_events_notes_personal_id
		)
		VALUES (?,?)
	`,
		data.Path,
		data.SchoolClassEventsNotesPersonalData.ID,
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

func (s *SchoolClassEventsNotesPersonalPicturesMySQL) GetByClassEventsNotesID(id uint) ([]school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesData, error) {
	rows, err := s.db.Query(`
		SELECT
			cenp.id,
			cenp.path,
			cenp.class_events_notes_personal_id
		FROM class_events_notes_personal_pictures AS cenp
		WHERE cenp.class_events_notes_personal_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var notes []school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesData
	for rows.Next() {
		var note school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesData
		rows.Scan(
			&note.ID,
			&note.Path,
			&note.SchoolClassEventsNotesPersonalData.ID,
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
	school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesDB = &SchoolClassEventsNotesPersonalPicturesMySQL{
		db: db,
	}
}
