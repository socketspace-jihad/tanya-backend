package student_presensi

type StudentPresensiRepository interface {
	GetByID(uint) (*StudentPresensiData, error)
	Save(*StudentPresensiData) error
	GetByStudentProfilesID(uint) ([]StudentPresensiData, error)
}
