package student_profiles_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type StudentProfilesMySQL struct {
	db *sql.DB
}

func NewStudentProfilesMySQL() student_profiles.StudentProfilesRepo {
	return &StudentProfilesMySQL{}
}

func (s *StudentProfilesMySQL) GetByID(id uint) (*student_profiles.StudentProfilesData, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("SELECT id, school_id, nisn,current_school_class_id FROM student_profiles WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	var data student_profiles.StudentProfilesData
	for rows.Next() {
		rows.Scan(&data.ID, &data.SchoolData.ID, &data.NISN, &data.CurrentSchoolData.ID)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *StudentProfilesMySQL) GetByUserRoleID(userRoleId uint) (*student_profiles.StudentProfilesData, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("SELECT id, school_id, nisn, current_school_class_id,name FROM student_profiles WHERE user_roles_id=?", userRoleId)
	if err != nil {
		return nil, err
	}
	var data student_profiles.StudentProfilesData
	for rows.Next() {
		rows.Scan(
			&data.ID,
			&data.SchoolData.ID,
			&data.NISN,
			&data.CurrentSchoolData.ID,
			&data.Name,
		)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *StudentProfilesMySQL) Save(data *student_profiles.StudentProfilesData) error {
	return nil
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
	student_profiles.StudentProfilesDB = &StudentProfilesMySQL{
		db: db,
	}
}
