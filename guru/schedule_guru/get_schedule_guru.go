package schedule_guru

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
)

type ScheduleGuru struct {
	SchoolClassEvents []school_class_events.SchoolClassEventsData `json:"school_class_events"`
}

func (s *ScheduleGuru) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println(g)
	s.SchoolClassEvents, err = school_class_events.SchoolClassEventDB.GetByTeacherProfilesID(g.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/guru/schedule", middlewares.TeacherMiddleware(
		&ScheduleGuru{}),
	)
}
