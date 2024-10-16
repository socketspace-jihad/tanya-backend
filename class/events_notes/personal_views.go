package events_notes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_viewer"
)

type CatatanPersonalViews struct{}

func (c *CatatanPersonalViews) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query().Get("class_events_id")
		if params == "" {
			http.Error(w, "class_events_id must be defined as params", http.StatusBadRequest)
			return
		}
		classEventsID, err := strconv.Atoi(params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		viewer, err := school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerDB.GetByClassEventsNotesID(uint(classEventsID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(viewer)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/class/events/notes/personal/viewer", auth.AuthMiddlewareHandler(&CatatanPersonalViews{}))
}
