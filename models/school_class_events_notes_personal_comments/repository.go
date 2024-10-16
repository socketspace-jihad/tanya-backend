package school_class_events_notes_personal_comments

type SchoolClassEventsNotesPersonalCommentsRepository interface {
	Save(*SchoolClassEventsNotesPersonalCommentsData) error
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesPersonalCommentsData, error)
}
