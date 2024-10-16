package school_class_events_notes_personal_viewer_mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_viewer"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalViewerMySQL struct {
	*sql.DB
}

func (s *SchoolClassEventsNotesPersonalViewerMySQL) Save(data *school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData) error {
	tx, err := s.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	var teacherProfilesID sql.NullInt64
	if data.TeacherProfilesData != nil {
		teacherProfilesID = sql.NullInt64{
			Int64: int64(data.TeacherProfilesData.ID),
			Valid: true,
		}
	}

	var parentProfilesID sql.NullInt64
	if data.ParentProfilesData != nil {
		parentProfilesID = sql.NullInt64{
			Int64: int64(data.ParentProfilesData.ID),
			Valid: true,
		}
	}

	var studentProfilesID sql.NullInt64
	if data.StudentProfilesData != nil {
		studentProfilesID = sql.NullInt64{
			Int64: int64(data.StudentProfilesData.ID),
			Valid: true,
		}
	}

	res, err := tx.Exec(`
		INSERT IGNORE INTO
			class_events_notes_personal_viewers
		(
			class_events_notes_personal_id,
			teacher_profiles_id,
			parent_profiles_id,
			student_profiles_id	
		) VALUES (?,?,?,?)
	`,
		data.SchoolClassEventsNotesData.ID,
		teacherProfilesID,
		parentProfilesID,
		studentProfilesID,
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
	tx.Commit()
	data.ID = uint(id)
	return nil
}

func (s *SchoolClassEventsNotesPersonalViewerMySQL) GetByID(uint) (*school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData, error) {
	return nil, nil
}

func (s *SchoolClassEventsNotesPersonalViewerMySQL) GetByClassEventsNotesID(id uint) ([]school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData, error) {
	rows, err := s.DB.Query(`
		SELECT
			cenv.id,
			IFNULL(tp.id,0),
			IFNULL(tp.name,""),
			IFNULL(pp.id,0),
			IFNULL(pp.name,""),
			IFNULL(sp.id,0),
			IFNULL(sp.name,"")
		FROM class_events_notes_personal_viewers AS cenv
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = cenv.teacher_profiles_id
		LEFT JOIN parent_profiles AS pp
			ON pp.id = cenv.parent_profiles_id
		LEFT JOIN student_profiles AS sp
			ON sp.id = cenv.student_profiles_id
		WHERE cenv.class_events_notes_personal_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var events []school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData
	for rows.Next() {
		event := school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
			ParentProfilesData:  &parent_profiles.ParentProfilesData{},
			StudentProfilesData: &student_profiles.StudentProfilesData{},
		}
		if err := rows.Scan(
			&event.ID,
			&event.TeacherProfilesData.ID,
			&event.TeacherProfilesData.Name,
			&event.ParentProfilesData.ID,
			&event.ParentProfilesData.Name,
			&event.StudentProfilesData.ID,
			&event.StudentProfilesData.Name,
		); err != nil {
			log.Println("ERR", err)
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
	school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerDB = &SchoolClassEventsNotesPersonalViewerMySQL{
		DB: db,
	}
}
