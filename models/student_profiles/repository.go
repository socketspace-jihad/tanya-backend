package student_profiles

type StudentProfilesRepo interface {
	GetByID(uint) (*StudentProfilesData, error)
	GetByUserRoleID(uint) (*StudentProfilesData, error)
	Save(*StudentProfilesData) error
}
