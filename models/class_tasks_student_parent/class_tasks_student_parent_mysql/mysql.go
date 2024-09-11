package class_tasks_student_parent_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student_parent"
)

type ClassTasksStudentParentMySQL struct {
	*sql.DB
}

func (c *ClassTasksStudentParentMySQL) Save(data *class_tasks_student_parent.ClassTasksStudentParentData) error {
	tx, err := c.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Exec(`
		INSERT INTO
			class_tasks_student_parent
		(
			class_tasks_student_id,
			parent_profiles_id,
			keterangan
		)
			VALUES
		(?,?,?)
	`, data.ClassTasksStudentData.ID, data.ParentProfilesData.ID, data.Keterangan); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *ClassTasksStudentParentMySQL) GetByID(id uint) (*class_tasks_student_parent.ClassTasksStudentParentData, error) {
	return nil, nil
}

func (c *ClassTasksStudentParentMySQL) GetByParentProfilesID(id uint) ([]class_tasks_student_parent.ClassTasksStudentParentData, error) {
	return nil, nil
}

func (c *ClassTasksStudentParentMySQL) GetByClassTasksStudentID(id uint) ([]class_tasks_student_parent.ClassTasksStudentParentData, error) {
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
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	class_tasks_student_parent.ClassTasksStudentParentDB = &ClassTasksStudentParentMySQL{
		DB: db,
	}
}
