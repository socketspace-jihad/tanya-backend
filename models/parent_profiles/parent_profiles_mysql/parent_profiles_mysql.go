package parent_profiles_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type ParentProfilesMySQL struct {
	db *sql.DB
}

func (p *ParentProfilesMySQL) GetByID(id uint) (*parent_profiles.ParentProfilesData, error) {
	rows, err := p.db.Query(`
		SELECT
			pp.id,
			pp.name,
			u.first_name,
			u.email
		FROM
			parent_profiles AS pp
		LEFT JOIN user_roles AS ur
			ON ur.id = pp.user_roles_id
		LEFT JOIN user AS u
			ON u.id = ur.user_id
		WHERE pp.id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	data := parent_profiles.ParentProfilesData{
		UserRolesData: user_roles.UserRolesData{
			UserData: user.UserData{},
		},
	}
	for rows.Next() {
		if err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.UserRolesData.UserData.FirstName,
			&data.UserRolesData.UserData.Email,
		); err != nil {
			return nil, err
		}
	}
	return &data, nil
}

func (p *ParentProfilesMySQL) Save(data *parent_profiles.ParentProfilesData) error {
	tx, err := p.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(`
		INSERT INTO
			parent_profiles
		(
			user_roles_id,
			nik,
			name
		) VALUES (
			?,?,?
		) 
	`,
		data.UserRolesData.ID,
		data.NIK,
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

func (p *ParentProfilesMySQL) GetByUserRoleID(id uint) (*parent_profiles.ParentProfilesData, error) {
	q, err := p.db.Query("SELECT id,name FROM parent_profiles WHERE user_roles_id=?", id)
	if err != nil {
		return nil, err
	}
	data := &parent_profiles.ParentProfilesData{}
	for q.Next() {
		q.Scan(&data.ID, &data.Name)
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
	parent_profiles.ParentProfilesDB = &ParentProfilesMySQL{
		db: db,
	}
}
