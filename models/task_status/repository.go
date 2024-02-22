package task_status

type TaskStatusRepository interface {
	GetByID(uint) *TaskStatusData
}
