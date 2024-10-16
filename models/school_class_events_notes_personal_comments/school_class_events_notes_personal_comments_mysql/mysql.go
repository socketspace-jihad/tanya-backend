package school_class_events_notes_personal_comments_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_comments"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsNotesPersonalCommentsMySQL struct {
	db *sql.DB
}

func (s *SchoolClassEventsNotesPersonalCommentsMySQL) Save(data *school_class_events_notes_personal_comments.SchoolClassEventsNotesPersonalCommentsData) error {
	tx, err := s.db.Begin()
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
		INSERT INTO
			class_events_notes_personal_comments
		(
			class_events_notes_personal_id,
			teacher_profiles_id,
			parent_profiles_id,
			student_profiles_id,
			content
		)
		VALUES (?,?,?,?,?)
	`,
		data.SchoolClassEventsNotesData.ID,
		teacherProfilesID,
		parentProfilesID,
		studentProfilesID,
		data.Content,
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

func (s *SchoolClassEventsNotesPersonalCommentsMySQL) GetByClassEventsNotesID(id uint) ([]school_class_events_notes_personal_comments.SchoolClassEventsNotesPersonalCommentsData, error) {
	rows, err := s.db.Query(`
		SELECT
			cenc.id,
			IFNULL(tp.id,0),
			IFNULL(tp.name,""),
			IFNULL(pp.id,0),
			IFNULL(pp.name,""),
			IFNULL(sp.id,0),
			IFNULL(sp.name,""),
			content,
			cenc.created_at
		FROM class_events_notes_personal_comments AS cenc
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = cenc.teacher_profiles_id
		LEFT JOIN parent_profiles AS pp
			ON pp.id = cenc.parent_profiles_id
		LEFT JOIN student_profiles AS sp
			ON sp.id = cenc.student_profiles_id
		WHERE cenc.class_events_notes_personal_id = ?
		ORDER BY created_at DESC
	`, id)
	if err != nil {
		return nil, err
	}
	var comments []school_class_events_notes_personal_comments.SchoolClassEventsNotesPersonalCommentsData
	for rows.Next() {
		comment := school_class_events_notes_personal_comments.SchoolClassEventsNotesPersonalCommentsData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
			ParentProfilesData:  &parent_profiles.ParentProfilesData{},
			StudentProfilesData: &student_profiles.StudentProfilesData{},
		}
		rows.Scan(
			&comment.ID,
			&comment.TeacherProfilesData.ID,
			&comment.TeacherProfilesData.Name,
			&comment.ParentProfilesData.ID,
			&comment.ParentProfilesData.Name,
			&comment.StudentProfilesData.ID,
			&comment.StudentProfilesData.Name,
			&comment.Content,
			&comment.CreatedAt,
		)
		comments = append(comments, comment)
	}
	return comments, nil
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
	school_class_events_notes_personal_comments.SchoolClassEventsNotesPersonalCommentsDB = &SchoolClassEventsNotesPersonalCommentsMySQL{
		db: db,
	}
}
