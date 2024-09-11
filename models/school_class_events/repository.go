package school_class_events

import "time"

type SchoolClassEventRepo interface {
	Save(*SchoolClassEventsData) error
	GetByID(uint) (*SchoolClassEventsData, error)
	GetBySchoolClassID(uint) ([]SchoolClassEventsData, error)
	GetBySchoolClassIDAndTimestamp(uint, time.Time) ([]SchoolClassEventsData, error)
	GetByTeacherProfilesID(uint) ([]SchoolClassEventsData, error)
	GetByTeacherProfilesIDAndTimeRange(uint, time.Time) ([]SchoolClassEventsData, error)
	GetByStudentProfilesID(uint) ([]SchoolClassEventsData, error)
	GetNearestStudentEventsByTimeAndID(uint, uint, time.Time) (*SchoolClassEventsData, error)
}
