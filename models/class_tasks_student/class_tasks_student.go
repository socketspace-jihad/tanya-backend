package class_tasks_student

import (
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/task_status"
)

type ClassTasksStudentData struct {
	ID                                   uint `json:"id"`
	class_tasks.ClassTasksData           `json:"class_tasks"`
	student_profiles.StudentProfilesData `json:"student_profiles"`
	task_status.TaskStatusData           `json:"tasks_status"`
	Score                                *uint `json:"score"`
	ApproverCount                        uint  `json:"approver_count"`
}
