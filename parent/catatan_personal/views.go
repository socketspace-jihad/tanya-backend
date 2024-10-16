package catatan_personal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_viewer"
)

type CatatanPersonalView struct{}

func (c *CatatanPersonalView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		view := school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerData{}
		if err := json.Unmarshal(body, &view); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		view.ParentProfilesData = parent
		if err := school_class_events_notes_personal_viewer.SchoolClassEventsNotesPersonalViewerDB.Save(&view); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/parent/catatan-personal/view", middlewares.ParentMiddleware(&CatatanPersonalView{}))
}
