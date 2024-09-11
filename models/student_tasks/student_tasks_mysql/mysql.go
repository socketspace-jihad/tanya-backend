package student_tasks_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/student_tasks"
)

type StudentTasksMySQL struct {
	db *sql.DB
}

func (s *StudentTasksMySQL) GetByID(id uint) (*student_tasks.StudentTasksData, error) {
	return nil, nil
}

func (s *StudentTasksMySQL) GetByStudentID(id uint) ([]student_tasks.StudentTasksData, error) {
	rows, err := s.db.Query(`
		SELECT 
			st.id,
			st.judul,
			st.deskripsi,
			st.due_date,
			st.working_method_id,
			st.tasks_status_id,
			wm.name,
			subjects.name,
			tp.name,
			ts.name
		FROM student_tasks AS st
		LEFT JOIN working_method AS wm
			ON st.working_method_id = wm.id
		LEFT JOIN subjects
			ON st.subjects_id = subjects.id
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = st.teacher_profiles_id
		LEFT JOIN tasks_status AS ts
			ON ts.id = st.tasks_status_id
		WHERE student_profiles_id=?`, id)
	if err != nil {
		return nil, err
	}
	data := []student_tasks.StudentTasksData{}
	for rows.Next() {
		var task student_tasks.StudentTasksData
		rows.Scan(
			&task.ID,
			&task.Judul,
			&task.Deskripsi,
			&task.DueDate,
			&task.WorkingMethodData.ID,
			&task.TaskStatusData.ID,
			&task.WorkingMethodData.Name,
			&task.SubjectsData.Name,
			&task.TeacherProfilesData.Name,
			&task.TaskStatusData.Name,
		)
		data = append(data, task)
	}
	return data, nil
}

func (s *StudentTasksMySQL) GetByTasksID(id uint) (*student_tasks.StudentTasksData, error) {
	return nil, nil
}

func (s *StudentTasksMySQL) Save(data *student_tasks.StudentTasksData) error {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Query(
		`INSERT INTO student_tasks(
			student_profiles_id,
			teacher_profiles_id,
			subjects_id,
			judul,
			deskripsi,
			due_date,
			working_method_id
		) VALUES (?,?,?,?,?,?,?)`,
		data.StudentProfilesData.ID,
		data.TeacherProfilesData.ID,
		data.SubjectsData.ID,
		data.Judul,
		data.Deskripsi,
		data.DueDate,
		data.WorkingMethodData.ID,
	); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
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
	student_tasks.StudentTasksDB = &StudentTasksMySQL{
		db: db,
	}
}
