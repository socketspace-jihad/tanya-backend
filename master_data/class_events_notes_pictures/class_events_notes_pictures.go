package class_events_notes_pictures

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_pictures"
)

type ClassEventsNotesPictures struct{}

func (g *ClassEventsNotesPictures) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query().Get("school_class_events_notes_id")
	if params == "" {
		http.Error(w, "school_class_events_notes_id should be defined on query params", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pictures, err := school_class_events_notes_pictures.SchoolClassEventsNotesPicturesDB.GetByClassEventsNotesID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pictures)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/master-data/class-events-notes-pictures",
		auth.AuthMiddlewareHandler(&ClassEventsNotesPictures{}),
	)
}
