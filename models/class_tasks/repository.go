package class_tasks

type ClassTasksRepository interface {
	GetByID(uint) (*ClassTasksData, error)
	GetByTeacherProfilesID(uint) ([]ClassTasksData, error)
	GetBySchoolClassID(uint) ([]ClassTasksData, error)
	Save(*ClassTasksData) error
}
