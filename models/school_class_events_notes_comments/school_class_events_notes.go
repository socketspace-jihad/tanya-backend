package school_class_events_notes_comments

import "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"

type SchoolClassEventsNotesCommentsData struct {
	ID                                                   uint   `json:"id"`
	Path                                                 string `json:"path"`
	school_class_events_notes.SchoolClassEventsNotesData `json:"school_class_events_notes"`
}
