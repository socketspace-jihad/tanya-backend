package school_class_events_mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/event_types"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type SchoolClassEventsMySQL struct {
	db *sql.DB
}

func (s *SchoolClassEventsMySQL) Save(data *school_class_events.SchoolClassEventsData) error {
	return nil
}

func (s *SchoolClassEventsMySQL) GetByID(id uint) (*school_class_events.SchoolClassEventsData, error) {
	return nil, nil
}

func (s *SchoolClassEventsMySQL) GetByTeacherProfilesID(id uint) ([]school_class_events.SchoolClassEventsData, error) {
	row, err := s.db.Query(`
		SELECT
			ce.id,
			ce.name,
			ce.start_date,
			ce.end_date,
			sc.name,
			sr.name,
			tp.name,
			sc.id
		FROM
			class_events AS ce
		LEFT JOIN school_class AS sc
			ON sc.id = ce.school_class_id
		LEFT JOIN school_rooms AS sr
			ON sr.id =  ce.school_rooms_id
		LEFT JOIN teacher_profiles AS tp
			ON ce.teacher_profiles_id = tp.id
		WHERE ce.teacher_profiles_id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	var events []school_class_events.SchoolClassEventsData

	for row.Next() {
		data := school_class_events.SchoolClassEventsData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		}
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SchoolClassData.Name,
			&data.SchoolRoomsData.Name,
			&data.TeacherProfilesData.Name,
			&data.SchoolClassData.ID,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
}

func (s *SchoolClassEventsMySQL) GetByTeacherProfilesIDAndTimeRange(id uint, time time.Time) ([]school_class_events.SchoolClassEventsData, error) {
	row, err := s.db.Query(`
		SELECT
			ce.id,
			ce.name,
			ce.start_date,
			ce.end_date,
			sc.name,
			sr.name,
			tp.name,
			sc.id
		FROM
			class_events AS ce
		LEFT JOIN school_class AS sc
			ON sc.id = ce.school_class_id
		LEFT JOIN school_rooms AS sr
			ON sr.id =  ce.school_rooms_id
		LEFT JOIN teacher_profiles AS tp
			ON ce.teacher_profiles_id = tp.id
		WHERE ce.teacher_profiles_id = ? AND YEAR(ce.start_date) = YEAR(?) AND MONTH(ce.start_date) = MONTH(?) AND DAY(ce.start_date) = DAY(?)
	`, id, time, time, time)
	if err != nil {
		return nil, err
	}

	var events []school_class_events.SchoolClassEventsData

	for row.Next() {
		data := school_class_events.SchoolClassEventsData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		}
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SchoolClassData.Name,
			&data.SchoolRoomsData.Name,
			&data.TeacherProfilesData.Name,
			&data.SchoolClassData.ID,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
}

func (s *SchoolClassEventsMySQL) GetBySchoolClassID(id uint) ([]school_class_events.SchoolClassEventsData, error) {
	row, err := s.db.Query(`
		SELECT
			ce.id,
			ce.name,
			ce.start_date,
			ce.end_date,
			sc.name,
			sr.name,
			tp.name,
			s.name
		FROM
			class_events AS ce
		LEFT JOIN school_class AS sc
			ON sc.id = ce.school_class_id
		LEFT JOIN schools AS s
			ON s.id = sc.school_id
		LEFT JOIN school_rooms AS sr
			ON sr.id =  ce.school_rooms_id
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = ce.teacher_profiles_id
		WHERE ce.school_class_id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	var events []school_class_events.SchoolClassEventsData

	for row.Next() {
		data := school_class_events.SchoolClassEventsData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		}
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SchoolClassData.Name,
			&data.SchoolRoomsData.Name,
			&data.TeacherProfilesData.Name,
			&data.SchoolClassData.SchoolData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
}

func (s *SchoolClassEventsMySQL) GetBySchoolClassIDAndTimestamp(id uint, t time.Time) ([]school_class_events.SchoolClassEventsData, error) {
	row, err := s.db.Query(`
	SELECT
		ce.id,
		ce.name,
		ce.start_date,
		ce.end_date,
		sc.name,
		sr.name,
		tp.name
	FROM
		class_events AS ce
	LEFT JOIN school_class AS sc
		ON sc.id = ce.school_class_id
	LEFT JOIN school_rooms AS sr
		ON sr.id =  ce.school_rooms_id
	LEFT JOIN teacher_profiles AS tp
		ON tp.id = ce.teacher_profiles_id
	WHERE ce.school_class_id = ? AND DAY(ce.start_date) = DAY(?) AND MONTH(ce.start_date) = MONTH(?) AND YEAR(ce.start_date) = YEAR(?)
`, id, t, t, t)
	if err != nil {
		return nil, err
	}

	var events []school_class_events.SchoolClassEventsData

	for row.Next() {
		data := school_class_events.SchoolClassEventsData{
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		}
		if err := row.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SchoolClassData.Name,
			&data.SchoolRoomsData.Name,
			&data.TeacherProfilesData.Name,
		); err != nil {
			return nil, err
		}
		events = append(events, data)
	}

	return events, nil
}

func (s *SchoolClassEventsMySQL) GetByStudentProfilesID(id uint) ([]school_class_events.SchoolClassEventsData, error) {
	return nil, nil
}

func (s *SchoolClassEventsMySQL) GetNearestStudentEventsByTimeAndID(studentID uint, schoolClassID uint, t time.Time) (*school_class_events.SchoolClassEventsData, error) {
	t = t.Add(-7 * time.Hour)
	rows, err := s.db.Query(`
	SELECT
		ce.id,
		ce.name,
		ce.start_date,
		ce.end_date,
		sc.name,
		sr.name,
		tp.name,
		sp.created_at AS presensi_date,
		s.name
	FROM
		class_events AS ce
	LEFT JOIN school_class AS sc
		ON sc.id = ce.school_class_id
	LEFT JOIN subjects AS s
		ON s.id = ce.subjects_id
	LEFT JOIN school_rooms AS sr
		ON sr.id =  ce.school_rooms_id
	LEFT JOIN teacher_profiles AS tp
		ON tp.id = ce.teacher_profiles_id
	LEFT JOIN student_presensi AS sp
		ON sp.events_id = ce.id AND sp.event_types_id = ? AND student_profiles_id = ?
	WHERE ce.school_class_id = ? AND ce.end_date > ? AND DAY(?) = DAY(ce.end_date)
	ORDER BY ce.end_date
	LIMIT 1
	`, event_types.ClassEvents.ID, studentID, schoolClassID, t, t)
	if err != nil {
		return nil, err
	}
	data := &school_class_events.SchoolClassEventsData{
		TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		SubjectsData:        &subjects.SubjectsData{},
	}
	for rows.Next() {
		if err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.StartDate,
			&data.EndDate,
			&data.SchoolClassData.Name,
			&data.SchoolRoomsData.Name,
			&data.TeacherProfilesData.Name,
			&data.PresensiDate,
			&data.SubjectsData.Name,
		); err != nil {
			return nil, err
		}
	}
	return data, nil
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
	school_class_events.SchoolClassEventDB = &SchoolClassEventsMySQL{
		db: db,
	}
}
