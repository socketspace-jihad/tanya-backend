package events_notes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"
)

type EventsNotes struct{}

func (e *EventsNotes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("class_events_id")
	if data == "" {
		return
	}
	classEventsID, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	notes, err := school_class_events_notes.SchoolClassEventsNotesDB.GetByClassEventsID(uint(classEventsID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func init() {
	http.DefaultServeMux.Handle("/v1/class/events/notes", &EventsNotes{})
}
