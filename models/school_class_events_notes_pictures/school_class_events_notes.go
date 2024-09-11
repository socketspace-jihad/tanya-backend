package school_class_events_notes_pictures

import "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"

type SchoolClassEventsNotesPicturesData struct {
	ID                                                   uint   `json:"id"`
	Path                                                 string `json:"path"`
	school_class_events_notes.SchoolClassEventsNotesData `json:"school_class_events_notes"`
}
