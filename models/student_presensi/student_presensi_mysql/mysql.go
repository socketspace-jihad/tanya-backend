package student_presensi_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
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
	student_presensi.StudentPresensiDB = &StudentPresensiMySQL{
		db: db,
	}
}
