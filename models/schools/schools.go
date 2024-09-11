package schools

import (
	"github.com/socketspace-jihad/tanya-backend/models/city"
	"github.com/socketspace-jihad/tanya-backend/models/district"
	"github.com/socketspace-jihad/tanya-backend/models/province"
	"github.com/socketspace-jihad/tanya-backend/models/school_levels"
)

type SchoolData struct {
	ID                             uint   `json:"id"`
	Name                           string `json:"name"`
	province.ProvinceData          `json:"province"`
	city.CityData                  `json:"city"`
	district.DistrictData          `json:"district"`
	school_levels.SchoolLevelsData `json:"school_levels"`
}
