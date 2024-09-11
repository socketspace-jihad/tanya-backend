package school_class

import (
	"github.com/socketspace-jihad/tanya-backend/models/school_levels"
	"github.com/socketspace-jihad/tanya-backend/models/schools"
)

type SchoolClassData struct {
	ID                             uint   `json:"id"`
	Name                           string `json:"name"`
	school_levels.SchoolLevelsData `json:"school_levels"`
	schools.SchoolData             `json:"schools"`
}
