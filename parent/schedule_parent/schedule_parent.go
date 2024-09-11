package schedule_parent

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/parent_student"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
)

type ScheduleParent struct {
	SchoolClassEvents []school_class_events.SchoolClassEventsData `json:"school_class_events"`
	mtx               *sync.Mutex
}

func (s *ScheduleParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query().Get("time")
	if params == "" {
		log.Println("ERROR TIME PARAMS")
		http.Error(w, "'time' query params must be defined", http.StatusBadRequest)
		return
	}
	t, err := time.Parse("2006-01-02", params)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students, err := parent_student.ParentStudentDB.GetStudentsByParentID(parent.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	studentsEvents := make(map[string][]school_class_events.SchoolClassEventsData)
	for _, student := range students {
		studentsEvents[student.Name], err = school_class_events.SchoolClassEventDB.GetBySchoolClassIDAndTimestamp(
			student.CurrentSchoolData.ID,
			t,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentsEvents)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/schedule", middlewares.ParentMiddleware(
		&ScheduleParent{
			mtx: &sync.Mutex{},
		}),
	)
}
