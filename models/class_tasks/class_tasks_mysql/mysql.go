package class_tasks_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks"
)

type ClassTasksMySQL struct {
	*sql.DB
}

func (c *ClassTasksMySQL) Save(data *class_tasks.ClassTasksData) error {
	return nil
}

func (c *ClassTasksMySQL) GetByID(id uint) (*class_tasks.ClassTasksData, error) {
	return nil, nil
}

func (c *ClassTasksMySQL) GetByTeacherProfilesID(id uint) ([]class_tasks.ClassTasksData, error) {
	return nil, nil
}

func (c *ClassTasksMySQL) GetBySchoolClassID(id uint) ([]class_tasks.ClassTasksData, error) {
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
	class_tasks.ClassTasksDB = &ClassTasksMySQL{
		DB: db,
	}
}
