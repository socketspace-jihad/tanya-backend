package class_tasks_student

type ClassTasksStudentRepository interface {
	GetByID(uint) (*ClassTasksStudentData, error)
	GetByStudentProfilesID(uint) ([]ClassTasksStudentData, error)
	GetByParentProfilesID(uint) ([]ClassTasksStudentData, error)
	Save(*ClassTasksStudentData) error
}
