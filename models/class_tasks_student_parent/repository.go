package class_tasks_student_parent

type ClassTasksStudentParent interface {
	Save(*ClassTasksStudentParentData) error
	GetByID(uint) (*ClassTasksStudentParentData, error)
	GetByClassTasksStudentID(uint) ([]ClassTasksStudentParentData, error)
	GetByParentProfilesID(uint) ([]ClassTasksStudentParentData, error)
}
