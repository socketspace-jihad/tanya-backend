package notification

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type NotificationData struct {
	ID                                   uint      `json:"id"`
	Title                                string    `json:"title"`
	Contents                             string    `json:"contents"`
	CreatedAt                            time.Time `json:"created_at"`
	TargetPath                           string    `json:"target_path"`
	ReadStatus                           bool      `json:"read_status"`
	student_profiles.StudentProfilesData `json:"student_profiles,omitempty"`
	teacher_profiles.TeacherProfilesData `json:"teacher_profiles,omitempty"`
	parent_profiles.ParentProfilesData   `json:"parent_profiles,omitempty"`
	user.UserData                        `json:"user,omitempty"`
}
