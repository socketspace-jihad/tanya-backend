package schedule_parent

import (
	"encoding/json"
	"net/http"
	"sync"

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
	students, err := parent_student.ParentStudentDB.GetStudentsByParentID(parent.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	studentsEvents := make(map[string][]school_class_events.SchoolClassEventsData)
	for _, student := range students {
		studentsEvents[student.Name], err = school_class_events.SchoolClassEventDB.GetBySchoolClassID(student.CurrentSchoolData.ID)
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
