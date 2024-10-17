package events

import "github.com/socketspace-jihad/tanya-backend/models/notification"

type PembuatanTugas struct {
	ClassID   uint
	TeacherID uint
	StudentID uint
	Message   string
	SubjectID uint8
	notification.NotificationData
}
