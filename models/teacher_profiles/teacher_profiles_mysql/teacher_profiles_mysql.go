package teacher_profiles_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type TeacherProfileMySQL struct {
	db *sql.DB
}

func (t *TeacherProfileMySQL) Save(profile *teacher_profiles.TeacherProfilesData) error {
	tx, err := t.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			teacher_profiles
		(
			user_roles_id,
			school_id,
			nuptk,
			name
		) VALUES (
			?,?,?,?
		) 
	`,
		profile.UserRolesData.ID,
		profile.SchoolData.ID,
		profile.NUPTK,
		profile.Name,
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
	profile.ID = uint(lastId)
	tx.Commit()
	return nil
}

func (t *TeacherProfileMySQL) GetByID(id uint) (*teacher_profiles.TeacherProfilesData, error) {
	q, err := t.db.Query(`
		SELECT 
			id,
			school_id,
			name,
			contact,
			address
		FROM teacher_profiles WHERE id=?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	data := &teacher_profiles.TeacherProfilesData{}
	for q.Next() {
		q.Scan(
			&data.ID,
			&data.SchoolData.ID,
			&data.Name,
			&data.Contact,
			&data.Address,
		)
	}
	return data, nil
}

func (t *TeacherProfileMySQL) GetByUserRoleID(userRoleID uint) (*teacher_profiles.TeacherProfilesData, error) {
	q, err := t.db.Query("SELECT id,school_id,name FROM teacher_profiles WHERE user_roles_id=?", userRoleID)
	if err != nil {
		return nil, err
	}
	data := &teacher_profiles.TeacherProfilesData{}
	for q.Next() {
		q.Scan(&data.ID, &data.SchoolData.ID, &data.Name)
	}
	return data, nil
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
	teacher_profiles.TeacherProfilesDB = &TeacherProfileMySQL{
		db: db,
	}
}
