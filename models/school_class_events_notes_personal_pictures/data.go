package school_class_events_notes_personal_pictures

import (
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal"
)

type SchoolClassEventsNotesPersonalPicturesData struct {
	ID                                                                    uint   `json:"id"`
	Path                                                                  string `json:"path"`
	school_class_events_notes_personal.SchoolClassEventsNotesPersonalData `json:"school_class_events_notes_personal"`
}
