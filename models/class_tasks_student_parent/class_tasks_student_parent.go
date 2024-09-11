package class_tasks_student_parent

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
)

type ClassTasksStudentParentData struct {
	ID                                        uint       `json:"id"`
	CreatedAt                                 *time.Time `json:"created_at"`
	parent_profiles.ParentProfilesData        `json:"parent_profiles"`
	class_tasks_student.ClassTasksStudentData `json:"class_tasks_student"`
	Keterangan                                string `json:"keterangan"`
}
