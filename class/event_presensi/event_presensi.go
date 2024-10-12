package event_presensi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
)

type StudentClass struct{}

func (e *StudentClass) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		data := r.URL.Query().Get("school_class_event_id")
		if data == "" {
			http.Error(w, "school_class_event_id must be defined", http.StatusBadRequest)
			return
		}
		eventID, err := strconv.Atoi(data)
		if err != nil {
			http.Error(w, "school_class_event_id params not valid", http.StatusBadRequest)
			return
		}
		presensi, err := student_presensi.StudentPresensiDB.GetBySchoolClassEventsID(uint(eventID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(presensi)

	case http.MethodPost:
		return
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/class/presensi", auth.AuthMiddlewareHandler(&StudentClass{}))
}
