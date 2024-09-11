package schedule_parent

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/utils"
)

type EventsByTime struct {
}

func (e *EventsByTime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	q = r.URL.Query().Get("student_profiles_id")
	studentID, err := strconv.Atoi(q)
	if err != nil {
		http.Error(w, "student_profiles_id not valid", http.StatusBadRequest)
		return
	}
	events, err := school_class_events.SchoolClassEventDB.GetNearestStudentEventsByTimeAndID(uint(studentID), class, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/parent/siswa/events/time",
		middlewares.ParentMiddleware(middlewares.StudentClassMiddlewares((&EventsByTime{}).ServeHTTP)),
	)
}
