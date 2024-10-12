package catatan_personal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal"
)

type CatatanPersonal struct {
}

func (c *CatatanPersonal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		data := r.URL.Query().Get("class_events_id")
		if data == "" {
			return
		}
		classEventsID, err := strconv.Atoi(data)
		if err != nil {
			return
		}
		notes, err := school_class_events_notes_personal.SchoolClassEventsNotesPersonalDB.GetByParentAndClassEventsID(uint(classEventsID), parent.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	default:
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/parent/catatan-personal", middlewares.ParentMiddleware(&CatatanPersonal{}))
}
