package school_class_events_notes_personal

type SchoolClassEventsNotesPersonalRepository interface {
	Save(*SchoolClassEventsNotesPersonalData) error
	GetByID(uint) (*SchoolClassEventsNotesPersonalData, error)
	GetByTeacherAndClassEventsID(uint, uint) ([]SchoolClassEventsNotesPersonalData, error)
	GetByParentAndClassEventsID(uint, uint) ([]SchoolClassEventsNotesPersonalData, error)
}
