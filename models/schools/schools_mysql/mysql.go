package schools_mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/socketspace-jihad/tanya-backend/models"
	"github.com/socketspace-jihad/tanya-backend/models/schools"
)

type SchoolsMySQL struct {
	db *sql.DB
}

func (s *SchoolsMySQL) GetByID(id uint) (*schools.SchoolData, error) {
	row, err := s.db.Query(`
		SELECT
			s.id,
			s.name,
			sl.name,
			p.id,
			p.name,
			c.id,
			c.name,
			d.id,
			d.name
		FROM schools AS s
		LEFT JOIN school_levels AS sl
			ON sl.id = s.school_levels_id
		LEFT JOIN province AS p
			ON p.id = s.province_id
		LEFT JOIN city AS c
			ON c.id = s.city_id
		LEFT JOIN district AS d
			ON d.id = s.district_id
		WHERE s.id = ?	
	`, id)
	if err != nil {
		return nil, err
	}
	school := &schools.SchoolData{}
	for row.Next() {
		if err := row.Scan(
			&school.ID,
			&school.Name,
			&school.SchoolLevelsData.Name,
			&school.ProvinceData.ID,
			&school.ProvinceData.Name,
			&school.CityData.ID,
			&school.CityData.Name,
			&school.DistrictData.ID,
			&school.DistrictData.Name,
		); err != nil {
			return nil, err
		}
	}
	return school, nil
}

func (s *SchoolsMySQL) GetAll() ([]schools.SchoolData, error) {
	rows, err := s.db.Query(`
		SELECT
			s.id,
			s.name,
			sl.name,
			p.id,
			p.name,
			c.id,
			c.name,
			d.id,
			d.name
		FROM schools AS s
		LEFT JOIN school_levels AS sl
			ON sl.id = s.school_levels_id
		LEFT JOIN province AS p
			ON p.id = s.province_id
		LEFT JOIN city AS c
			ON c.id = s.city_id
		LEFT JOIN district AS d
			ON d.id = s.district_id
	`)
	if err != nil {
		return nil, err
	}
	schoolsData := []schools.SchoolData{}
	for rows.Next() {
		school := schools.SchoolData{}
		if err := rows.Scan(
			&school.ID,
			&school.Name,
			&school.SchoolLevelsData.Name,
			&school.ProvinceData.ID,
			&school.ProvinceData.Name,
			&school.CityData.ID,
			&school.CityData.Name,
			&school.DistrictData.ID,
			&school.DistrictData.Name,
		); err != nil {
			return nil, err
		}
		schoolsData = append(schoolsData, school)
	}
	return schoolsData, nil
}

func (s *SchoolsMySQL) GetByNamePrefix(prefix string) ([]schools.SchoolData, error) {
	return nil, nil
}

func (s *SchoolsMySQL) GetByNPSNPrefix(prefix string) ([]schools.SchoolData, error) {
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
	schools.SchoolsDB = &SchoolsMySQL{
		db: db,
	}
}
