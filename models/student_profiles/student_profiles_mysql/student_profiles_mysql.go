package student_profiles_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rows, err := tx.Query(`
	SELECT
		sp.id,
		sp.school_id,
		sp.nisn,
		sp.current_school_class_id,
		sp.name,
		vs.id,
		vs.name,
		u.email
	FROM 
		student_profiles AS sp
	LEFT JOIN verified_status AS vs
		ON vs.id = sp.verified_status_id
	LEFT JOIN user_roles AS ur
		ON ur.id = sp.user_roles_id
	LEFT JOIN user AS u
		ON u.id = ur.user_id
	WHERE sp.id=?`, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var data student_profiles.StudentProfilesData
	for rows.Next() {
		if err := rows.Scan(
			&data.ID,
			&data.SchoolData.ID,
			&data.NISN,
			&data.CurrentSchoolData.ID,
			&data.Name,
			&data.VerifiedStatus.ID,
			&data.VerifiedStatus.Name,
			&data.UserRolesData.UserData.Email,
		); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &data, nil
}

func (s *StudentProfilesMySQL) GetByUserRoleID(userRoleId uint) (*student_profiles.StudentProfilesData, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rows, err := tx.Query(`
		SELECT
			sp.id,
			sp.school_id,
			sp.nisn,
			sp.current_school_class_id,
			sp.name,
			vs.id,
			vs.name,
			u.email,
			u.id
		FROM 
			student_profiles AS sp
		LEFT JOIN verified_status AS vs
			ON vs.id = sp.verified_status_id
		LEFT JOIN user_roles AS ur
			ON sp.user_roles_id = ur.id
		LEFT JOIN user AS u
			ON u.id = ur.user_id
		WHERE user_roles_id=?`, userRoleId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var data student_profiles.StudentProfilesData
	for rows.Next() {
		if err := rows.Scan(
			&data.ID,
			&data.SchoolData.ID,
			&data.NISN,
			&data.CurrentSchoolData.ID,
			&data.Name,
			&data.VerifiedStatus.ID,
			&data.VerifiedStatus.Name,
			&data.UserData.Email,
			&data.UserData.ID,
		); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &data, nil
}

func (s *StudentProfilesMySQL) Save(data *student_profiles.StudentProfilesData) error {
	tx, err := s.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil
	}
	res, err := tx.Exec(`
		INSERT INTO
			student_profiles
		(
			user_roles_id,
			school_id,
			nisn,
			name
		) VALUES (
			?,?,?,?
		) 
	`,
		data.UserRolesData.ID,
		data.SchoolData.ID,
		data.NISN,
		data.Name,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	data.ID = uint(lastId)
	tx.Commit()
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
