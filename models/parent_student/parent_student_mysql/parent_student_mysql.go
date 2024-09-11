package parent_student_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/parent_student"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type ParentStudentMySQL struct {
	db *sql.DB
}

func (p *ParentStudentMySQL) Save(data *parent_student.ParentStudentData) error {
	tx, err := p.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT IGNORE INTO
			parent_student
		(
			student_profiles_id,
			parent_profiles_id
		) VALUES (?,?)
	`, data.StudentProfilesData.ID, data.ParentProfilesData.ID)
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

func (p *ParentStudentMySQL) GetByID(id uint) (*parent_student.ParentStudentData, error) {
	return nil, nil
}

func (p *ParentStudentMySQL) GetParentsByStudentID(id uint) ([]parent_profiles.ParentProfilesData, error) {
	return nil, nil
}

func (p *ParentStudentMySQL) GetStudentsByParentID(id uint) ([]student_profiles.StudentProfilesData, error) {
	rows, err := p.db.Query(`
		SELECT
			sp.id,
			sp.current_school_class_id,
			sp.name,
			s.id,
			s.name,
			sc.id,
			sc.name
		FROM
			parent_student AS ps
		LEFT JOIN student_profiles AS sp
			ON sp.id = ps.student_profiles_id
		LEFT JOIN schools AS s
			ON s.id = sp.school_id
		LEFT JOIN school_class AS sc
			ON sc.id = sp.current_school_class_id
		WHERE ps.parent_profiles_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var students []student_profiles.StudentProfilesData
	for rows.Next() {
		var student student_profiles.StudentProfilesData
		if err := rows.Scan(
			&student.ID,
			&student.CurrentSchoolData.ID,
			&student.Name,
			&student.SchoolData.ID,
			&student.SchoolData.Name,
			&student.CurrentSchoolData.ID,
			&student.CurrentSchoolData.Name,
		); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
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
	parent_student.ParentStudentDB = &ParentStudentMySQL{
		db: db,
	}
}
