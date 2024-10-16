package catatan_personal

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal"
)

type CatatanPersonalList struct{}

func (c *CatatanPersonalList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teacher, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}
	switch r.Method {
	case http.MethodGet:
		notes, err := school_class_events_notes_personal.SchoolClassEventsNotesPersonalDB.GetByTeacherProfilesID(teacher.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/guru/catatan-personal/list", middlewares.TeacherMiddleware(&CatatanPersonalList{}))
}
