package class_tasks_student_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student"
)

type ClassTasksStudentMySQL struct {
	*sql.DB
}

func (c *ClassTasksStudentMySQL) GetByID(id uint) (*class_tasks_student.ClassTasksStudentData, error) {
	rows, err := c.DB.Query(`
		SELECT
			ct.id,
			ct.name,
			ct.judul,
			ct.deskripsi,
			ct.due_date,
			ct.created_at,
			cts.tasks_status_id,
			ts.name,
			cts.score,
			cts.created_at,
			tp.name
		FROM class_tasks_student AS cts
		LEFT JOIN class_tasks AS ct
			ON ct.id = cts.class_tasks_id
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = ct.teacher_profiles_id
		LEFT JOIN tasks_status AS ts
			ON ts.id = cts.tasks_status_id
		WHERE cts.id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var d class_tasks_student.ClassTasksStudentData
	for rows.Next() {
		if err := rows.Scan(
			&d.ClassTasksData.ID,
			&d.ClassTasksData.Name,
			&d.ClassTasksData.Judul,
			&d.ClassTasksData.Deskripsi,
			&d.ClassTasksData.DueDate,
			&d.ClassTasksData.CreatedAt,
			&d.TaskStatusData.ID,
			&d.TaskStatusData.Name,
			&d.Score,
			&d.CreatedAt,
			&d.TeacherProfilesData.Name,
		); err != nil {
			return nil, err
		}
	}
	return &d, nil
}

func (c *ClassTasksStudentMySQL) GetByStudentProfilesID(id uint) ([]class_tasks_student.ClassTasksStudentData, error) {
	rows, err := c.DB.Query(`
		SELECT
			ct.id,
			ct.name,
			ct.judul,
			ct.deskripsi,
			ct.due_date,
			ct.created_at,
			cts.tasks_status_id,
			ts.name,
			cts.score,
			cts.created_at,
			tp.name
		FROM class_tasks_student AS cts
		LEFT JOIN class_tasks AS ct
			ON ct.id = cts.class_tasks_id
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = ct.teacher_profiles_id
		LEFT JOIN tasks_status AS ts
			ON ts.id = cts.tasks_status_id
		WHERE cts.student_profiles_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	var data []class_tasks_student.ClassTasksStudentData
	for rows.Next() {
		var d class_tasks_student.ClassTasksStudentData
		if err := rows.Scan(
			&d.ClassTasksData.ID,
			&d.ClassTasksData.Name,
			&d.ClassTasksData.Judul,
			&d.ClassTasksData.Deskripsi,
			&d.ClassTasksData.DueDate,
			&d.ClassTasksData.CreatedAt,
			&d.TaskStatusData.ID,
			&d.TaskStatusData.Name,
			&d.Score,
			&d.CreatedAt,
			&d.TeacherProfilesData.Name,
		); err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func (c *ClassTasksStudentMySQL) GetByParentProfilesID(id uint) ([]class_tasks_student.ClassTasksStudentData, error) {
	rows, err := c.DB.Query(`
		SELECT 
			cts.id,
			ct.id,
			ct.name,
			ct.judul,
			ct.deskripsi,
			ct.due_date,
			ct.created_at,
			subjects.name,
			cts.tasks_status_id,
			ts.name,
			cts.score,
			cts.created_at,
			sp.name,
			sp.id,
			tp.name,
			sc.name,
			s.name,
			COUNT(ctsp.id) AS approver_count
		FROM parent_student AS ps
		RIGHT JOIN class_tasks_student AS cts
			ON cts.student_profiles_id  = ps.student_profiles_id
		LEFT JOIN student_profiles AS sp
			ON cts.student_profiles_id = sp.id
		LEFT JOIN tasks_status AS ts
			ON ts.id = cts.tasks_status_id
		LEFT JOIN class_tasks AS ct
			ON ct.id = cts.class_tasks_id AND ct.class_tasks_type_id = 2
		LEFT JOIN subjects
			ON subjects.id = ct.subject_id
		LEFT JOIN school_class AS sc
			ON sc.id = ct.school_class_id
		LEFT JOIN schools AS s
			ON s.id = sc.school_id
		LEFT JOIN teacher_profiles AS tp
			ON tp.id = ct.teacher_profiles_id
		LEFT JOIN class_tasks_student_parent  AS ctsp
			ON cts.id = ctsp.class_tasks_student_id
		WHERE ps.parent_profiles_id = ?
		GROUP BY cts.id 
		ORDER BY cts.created_at DESC;
	`, id)
	if err != nil {
		return nil, err
	}
	var data []class_tasks_student.ClassTasksStudentData
	for rows.Next() {
		var d class_tasks_student.ClassTasksStudentData
		if err := rows.Scan(
			&d.ID,
			&d.ClassTasksData.ID,
			&d.ClassTasksData.Name,
			&d.ClassTasksData.Judul,
			&d.ClassTasksData.Deskripsi,
			&d.ClassTasksData.DueDate,
			&d.ClassTasksData.CreatedAt,
			&d.ClassTasksData.SubjectsData.Name,
			&d.TaskStatusData.ID,
			&d.TaskStatusData.Name,
			&d.Score,
			&d.CreatedAt,
			&d.StudentProfilesData.Name,
			&d.StudentProfilesData.ID,
			&d.TeacherProfilesData.Name,
			&d.ClassTasksData.SchoolClassData.Name,
			&d.ClassTasksData.SchoolClassData.SchoolData.Name,
			&d.ApproverCount,
		); err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func (c *ClassTasksStudentMySQL) Save(data *class_tasks_student.ClassTasksStudentData) error {
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
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	class_tasks_student.ClassTasksStudentDB = &ClassTasksStudentMySQL{
		DB: db,
	}
}
