package school_class_events_notes_comments

type SchoolClassEventsNotesCommentsRepository interface {
	Save(*SchoolClassEventsNotesCommentsData) error
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesCommentsData, error)
}
