package student_events

import "time"

type StudentEventsRepository interface {
	GetByID(uint) (*StudentEventsData, error)
	GetByStudentID(uint) ([]StudentEventsData, error)
	GetByStudentIDAndTimestamp(uint, time.Time) ([]StudentEventsData, error)
	Save(*StudentEventsData) error
}
