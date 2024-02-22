package student_tasks

type StudentTasksRepository interface {
	Save(*StudentTasksData) error
	GetByID(uint) (*StudentTasksData, error)
	GetByStudentID(uint) ([]StudentTasksData, error)
	GetByTasksID(uint) (*StudentTasksData, error)
}
