package class_tasks

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_type"
	"github.com/socketspace-jihad/tanya-backend/models/school_class"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type ClassTasksData struct {
	ID                                   uint   `json:"id"`
	Name                                 string `json:"name"`
	class_tasks_type.ClassTasksTypeData  `json:"class_tasks_type"`
	subjects.SubjectsData                `json:"subjects"`
	teacher_profiles.TeacherProfilesData `json:"teacher_profiles"`
	school_class.SchoolClassData         `json:"school_class"`
	Judul                                string     `json:"judul"`
	Deskripsi                            string     `json:"deskripsi"`
	CreatedAt                            *time.Time `json:"created_at"`
	DueDate                              *time.Time `json:"due_date"`
	MaxAttempt                           uint8      `json:"max_attempt"`
}
