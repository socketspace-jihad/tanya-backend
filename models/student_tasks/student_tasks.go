package student_tasks

import (
	"github.com/socketspace-jihad/tanya-backend/models/deleted_status"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/subjects"
	"github.com/socketspace-jihad/tanya-backend/models/tasks_status"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/working_method"
)

type StudentTasksData struct {
	ID                                   uint `json:"id"`
	student_profiles.StudentProfilesData `json:"student_profiles"`
	teacher_profiles.TeacherProfilesData `json:"teacher_profiles"`
	subjects.SubjectsData                `json:"subjects"`
	deleted_status.DeletedStatusData     `json:"deleted_status"`
	Judul                                string `json:"judul"`
	Deskripsi                            string `json:"deskripsi"`
	DueDate                              string `json:"due_date"`
	working_method.WorkingMethodData     `json:"working_method"`
	CreatedAt                            string `json:"created_at"`
	tasks_status.TaskStatusData          `json:"tasks_status"`
}
