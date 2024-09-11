package schools

type SchoolRepo interface {
	GetByID(uint) (*SchoolData, error)
	GetAll() ([]SchoolData, error)
	GetByNamePrefix(string) ([]SchoolData, error)
	GetByNPSNPrefix(string) ([]SchoolData, error)
}
