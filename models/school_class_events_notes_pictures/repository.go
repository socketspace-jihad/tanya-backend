package school_class_events_notes_pictures

type SchoolClassEventsNotesPicturesRepository interface {
	Save(*SchoolClassEventsNotesPicturesData) error
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesPicturesData, error)
}
