package student_presensi_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/presensi_types"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
)

type StudentPresensiMySQL struct {
	db *sql.DB
}

func (s *StudentPresensiMySQL) Save(data *student_presensi.StudentPresensiData) error {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return tx.Rollback()
	}
	if _, err := tx.Query(`
			INSERT INTO student_presensi (
				student_profiles_id,
				events_id,
				event_types_id,
				lattitude,
				longitude
			) VALUES (
				?,?,?,?,?
			)
		`,
		data.StudentProfilesData.ID,
		data.EventsID,
		data.EventTypesData.ID,
		data.Lattitude,
		data.Longitude,
	); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *StudentPresensiMySQL) GetByID(id uint) (*student_presensi.StudentPresensiData, error) {
	return nil, nil
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
			pt.deskripsi
		FROM
			student_presensi AS sp
		LEFT JOIN presensi_types AS pt
			ON pt.id = sp.presensi_types_id
		LEFT JOIN class_events AS ce
			ON ce.id = sp.events_id AND sp.event_types_id = 3
		LEFT JOIN subjects AS s
			ON s.id = ce.subjects_id
		WHERE sp.student_profiles_id = ?
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
