package school_class_events_notes_viewer

type SchoolClassEventsNotesViewerRepo interface {
	Save(*SchoolClassEventsNotesViewerData) error
	GetByID(uint) (*SchoolClassEventsNotesViewerData, error)
	GetByClassEventsNotesID(uint) ([]SchoolClassEventsNotesViewerData, error)
}
