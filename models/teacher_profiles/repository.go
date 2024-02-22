package teacher_profiles

type TeacherProfilesRepo interface {
	GetByID(uint) (*TeacherProfilesData, error)
	GetByUserRoleID(uint) (*TeacherProfilesData, error)
	Save(*TeacherProfilesData) error
}
