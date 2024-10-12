package subjects

import "database/sql"

type SubjectsData struct {
	ID   uint           `json:"id"`
	Name sql.NullString `json:"name"`
}
