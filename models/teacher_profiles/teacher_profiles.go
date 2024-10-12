package teacher_profiles

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/schools"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
	"github.com/socketspace-jihad/tanya-backend/models/verified_status"
)

type TeacherProfilesData struct {
	ID                             uint   `json:"id",omitempty`
	Name                           string `json:"name"`
	user_roles.UserRolesData       `json:"user_roles"`
	schools.SchoolData             `json:"school"`
	subjects.SubjectsData          `json:"subjects",omitempty`
	JoinedAt                       time.Time `json:"joined_at"`
	CreatedAt                      time.Time `json:"created_at"`
	verified_status.VerifiedStatus `json:"verified_status"`
	NUPTK                          string `json:"nuptk"`
	Contact                        string `json:"contact"`
	Address                        string `json:"address"`
}
