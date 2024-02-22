package school_class_events_notes

type SchoolClassEventsNotesRepository interface {
	Save(*SchoolClassEventsNotesData) error
	GetByClassEventsID(uint) ([]SchoolClassEventsNotesData, error)
}
