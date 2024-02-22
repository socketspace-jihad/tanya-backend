package parent_profiles

type ParentProfilesRepository interface {
	GetByID(uint) (*ParentProfilesData, error)
	Save(*ParentProfilesData) error
	GetByUserRoleID(uint) (*ParentProfilesData, error)
}
