package schedule_siswa

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/utils"
)

type EventsByTime struct {
	*school_class_events.SchoolClassEventsData `json:"class_events"`
	*student_profiles.StudentProfilesData      `json:"student_profiles"`
}

func (e *EventsByTime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	student, err := middlewares.GetStudentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	class, err := middlewares.GetStudentClassIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	q := r.URL.Query().Get("time")
	if err != nil {
		http.Error(w, "time params cannot be nil", http.StatusBadRequest)
		return
	}
	t, err := time.Parse(utils.LayoutFormat, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := school_class_events.SchoolClassEventDB.GetNearestStudentEventsByTimeAndID(student.ID, class, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e.SchoolClassEventsData = events
	e.StudentProfilesData = student
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/siswa/events/time",
		middlewares.StudentMiddleware(middlewares.StudentClassMiddlewares((&EventsByTime{}).ServeHTTP)),
	)
}
