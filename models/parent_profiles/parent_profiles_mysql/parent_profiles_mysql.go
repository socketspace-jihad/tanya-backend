package parent_profiles_mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
)

type ParentProfilesMySQL struct {
	db *sql.DB
}

func (p *ParentProfilesMySQL) GetByID(id uint) (*parent_profiles.ParentProfilesData, error) {
	return nil, nil
}

func (p *ParentProfilesMySQL) Save(data *parent_profiles.ParentProfilesData) error {
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
