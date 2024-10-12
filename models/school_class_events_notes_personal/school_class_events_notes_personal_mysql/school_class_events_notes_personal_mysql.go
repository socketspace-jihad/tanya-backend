package school_class_events_notes_personal_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalMySQL struct {
	*sql.DB
}

func (s *SchoolClassEventsNotesPersonalMySQL) Save(data *school_class_events_notes_personal.SchoolClassEventsNotesPersonalData) error {
	tx, err := s.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			class_events_notes_personal
		(
			class_events_id,
			judul,
			deskripsi,
			teacher_profiles_id,
			student_profiles_id
		)
		VALUES (?,?,?,?,?)
	`,
		data.SchoolClassEventsData.ID,
		data.Judul,
		data.Deskripsi,
		data.TeacherProfilesData.ID,
		data.StudentProfilesData.ID,
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

func (s *SchoolClassEventsNotesPersonalMySQL) GetByID(uint) (*school_class_events_notes_personal.SchoolClassEventsNotesPersonalData, error) {
	return nil, nil
}

func (s *SchoolClassEventsNotesPersonalMySQL) GetByTeacherAndClassEventsID(classEvents uint, teacherID uint) ([]school_class_events_notes_personal.SchoolClassEventsNotesPersonalData, error) {
	rows, err := s.DB.Query(`
		SELECT
			cenp.id,
			cenp.judul,
			cenp.deskripsi,
			sp.id,
			sp.name
		FROM
			class_events_notes_personal AS cenp
		LEFT JOIN
			student_profiles AS sp
			ON sp.id = cenp.student_profiles_id
		WHERE
			cenp.teacher_profiles_id = ? AND cenp.class_events_id = ?
	`, teacherID, classEvents)
	if err != nil {
		return nil, err
	}
	var events []school_class_events_notes_personal.SchoolClassEventsNotesPersonalData
	for rows.Next() {
		event := school_class_events_notes_personal.SchoolClassEventsNotesPersonalData{
			StudentProfilesData: student_profiles.StudentProfilesData{},
		}
		if err := rows.Scan(
			&event.ID,
			&event.Judul,
			&event.Deskripsi,
			&event.StudentProfilesData.ID,
			&event.StudentProfilesData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (s *SchoolClassEventsNotesPersonalMySQL) GetByParentAndClassEventsID(classEvents uint, parentID uint) ([]school_class_events_notes_personal.SchoolClassEventsNotesPersonalData, error) {
	rows, err := s.DB.Query(`
		SELECT
			cenp.id,
			cenp.judul,
			cenp.deskripsi,
			sp.id,
			sp.name,
			tp.id,
			tp.name
		FROM
			class_events_notes_personal AS cenp
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = cenp.teacher_profiles_id
		LEFT JOIN
			student_profiles AS sp
			ON sp.id = cenp.student_profiles_id
		LEFT JOIN
			parent_student AS ps
			ON ps.student_profiles_id = sp.id
		WHERE
			ps.parent_profiles_id = ? AND cenp.class_events_id = ? 
	`, parentID, classEvents)
	if err != nil {
		return nil, err
	}
	var events []school_class_events_notes_personal.SchoolClassEventsNotesPersonalData
	for rows.Next() {
		event := school_class_events_notes_personal.SchoolClassEventsNotesPersonalData{
			StudentProfilesData: student_profiles.StudentProfilesData{},
			TeacherProfilesData: teacher_profiles.TeacherProfilesData{},
		}
		if err := rows.Scan(
			&event.ID,
			&event.Judul,
			&event.Deskripsi,
			&event.StudentProfilesData.ID,
			&event.StudentProfilesData.Name,
			&event.TeacherProfilesData.ID,
			&event.TeacherProfilesData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, event)
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
	school_class_events_notes_personal.SchoolClassEventsNotesPersonalDB = &SchoolClassEventsNotesPersonalMySQL{
		DB: db,
	}
}
