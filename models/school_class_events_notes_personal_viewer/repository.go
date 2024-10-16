package school_class_events_notes_personal_viewer

type SchoolClassEventsNotesPersonalViewerRepo interface {
	Save(*SchoolClassEventsNotesPersonalViewerData) error
	GetByID(uint) (*SchoolClassEventsNotesPersonalViewerData, error)
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesPersonalViewerData, error)
}
