package notification

type NotificationRepository interface {
	GetByID(uint) (*NotificationData, error)
	GetByUserID(uint) ([]NotificationData, error)
	GetByStudentProfilesID(uint) ([]NotificationData, error)
	GetByUserOrStudentProfilesID(uint, uint) ([]NotificationData, error)
}
