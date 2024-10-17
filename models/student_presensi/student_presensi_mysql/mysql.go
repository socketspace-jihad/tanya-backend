package student_presensi_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/presensi_types"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type StudentPresensiMySQL struct {
	db *sql.DB
}

func (s *StudentPresensiMySQL) Save(data *student_presensi.StudentPresensiData) error {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return tx.Rollback()
	}
	res, err := tx.Exec(`
			INSERT IGNORE INTO student_presensi (
				student_profiles_id,
				events_id,
				event_types_id,
				lattitude,
				longitude,
				class_events_id,
				teacher_profiles_id,
				hadir,
				presensi_types_id
			) VALUES (
				?,?,?,?,?,?,?,?,?
			)
		`,
		data.StudentProfilesData.ID,
		data.EventsID,
		data.EventTypesData.ID,
		data.Lattitude,
		data.Longitude,
		data.SchoolClassEventsData.ID,
		data.TeacherProfilesData.ID,
		data.Hadir,
		data.PresensitypesData.ID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	if affected == 0 {
		return errors.New("not rows affected")
	}
	return nil
}

func (s *StudentPresensiMySQL) GetByID(id uint) (*student_presensi.StudentPresensiData, error) {
	return nil, nil
}

func (s *StudentPresensiMySQL) GetBySchoolClassEventsID(id uint) ([]student_presensi.StudentPresensiData, error) {
	rows, err := s.db.Query(`
		SELECT
			sp.id,
			sp.created_at,
			ce.name,
			s.name,
			sp.lattitude,
			sp.longitude,
			sp.presensi_types_id,
			pt.name,
			pt.deskripsi,
			spr.name,
			sp.hadir
		FROM
			student_presensi AS sp
		LEFT JOIN presensi_types AS pt
			ON pt.id = sp.presensi_types_id
		LEFT JOIN class_events AS ce
			ON ce.id = sp.class_events_id
		LEFT JOIN subjects AS s
			ON s.id = ce.subjects_id
		LEFT JOIN student_profiles AS spr
			ON spr.id = sp.student_profiles_id
		WHERE sp.class_events_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	presensis := []student_presensi.StudentPresensiData{}
	for rows.Next() {
		presensi := student_presensi.StudentPresensiData{
			PresensitypesData: &presensi_types.PresensitypesData{},
			SchoolClassEventsData: &school_class_events.SchoolClassEventsData{
				SubjectsData: &subjects.SubjectsData{},
			},
			StudentProfilesData: &student_profiles.StudentProfilesData{},
		}
		if err := rows.Scan(
			&presensi.ID,
			&presensi.CreatedAt,
			&presensi.SchoolClassEventsData.Name,
			&presensi.SchoolClassEventsData.SubjectsData.Name,
			&presensi.Lattitude,
			&presensi.Longitude,
			&presensi.PresensitypesData.ID,
			&presensi.PresensitypesData.Name,
			&presensi.PresensitypesData.Deskripsi,
			&presensi.StudentProfilesData.Name,
			&presensi.Hadir,
		); err != nil {
			return nil, err
		}
		presensis = append(presensis, presensi)
	}
	return presensis, nil
}

func (s *StudentPresensiMySQL) GetByStudentProfilesID(id uint) ([]student_presensi.StudentPresensiData, error) {
	rows, err := s.db.Query(`
		SELECT
			sp.id,
			sp.created_at,
			ce.name,
			s.name,
			sp.lattitude,
			sp.longitude,
			sp.presensi_types_id,
			pt.name,
			pt.deskripsi,
			spr.name,
			sp.hadir,
			ce.start_date,
			ce.end_date,
			tp.id,
			tp.name,
			tp.contact
		FROM
			student_presensi AS sp
		LEFT JOIN presensi_types AS pt
			ON pt.id = sp.presensi_types_id
		LEFT JOIN class_events AS ce
			ON ce.id = sp.class_events_id
		LEFT JOIN subjects AS s
			ON s.id = ce.subjects_id
		LEFT JOIN student_profiles AS spr
			ON spr.id = sp.student_profiles_id
		LEFT JOIN teacher_profiles AS tp
			ON sp.teacher_profiles_id = tp.id
		WHERE sp.student_profiles_id = ?
		ORDER BY sp.created_at DESC
	`, id)
	if err != nil {
		return nil, err
	}
	var presensis []student_presensi.StudentPresensiData
	for rows.Next() {
		presensi := student_presensi.StudentPresensiData{
			PresensitypesData: &presensi_types.PresensitypesData{},
			SchoolClassEventsData: &school_class_events.SchoolClassEventsData{
				SubjectsData: &subjects.SubjectsData{},
			},
			StudentProfilesData: &student_profiles.StudentProfilesData{},
			TeacherProfilesData: &teacher_profiles.TeacherProfilesData{},
		}
		if err := rows.Scan(
			&presensi.ID,
			&presensi.CreatedAt,
			&presensi.SchoolClassEventsData.Name,
			&presensi.SchoolClassEventsData.SubjectsData.Name,
			&presensi.Lattitude,
			&presensi.Longitude,
			&presensi.PresensitypesData.ID,
			&presensi.PresensitypesData.Name,
			&presensi.PresensitypesData.Deskripsi,
			&presensi.StudentProfilesData.Name,
			&presensi.Hadir,
			&presensi.SchoolClassEventsData.StartDate,
			&presensi.SchoolClassEventsData.EndDate,
			&presensi.TeacherProfilesData.ID,
			&presensi.TeacherProfilesData.Name,
			&presensi.TeacherProfilesData.Contact,
		); err != nil {
			return nil, err
		}
		presensis = append(presensis, presensi)
	}
	return presensis, nil
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
	student_presensi.StudentPresensiDB = &StudentPresensiMySQL{
		db: db,
	}
}
