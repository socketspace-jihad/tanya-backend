package schools

type SchoolRepo interface {
	GetByID(uint) SchoolData
}
