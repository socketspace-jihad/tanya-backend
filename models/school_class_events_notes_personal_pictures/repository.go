package school_class_events_notes_personal_pictures

type SchoolClassEventsNotesPersonalPicturesRepository interface {
	Save(*SchoolClassEventsNotesPersonalPicturesData) error
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesPersonalPicturesData, error)
}
