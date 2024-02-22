package schedule_siswa

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_events"
	"github.com/socketspace-jihad/tanya-backend/utils"
)

type ScheduleSiswa struct {
	SchoolClassEvents []school_class_events.SchoolClassEventsData `json:"school_class_events"`
	StudentEvents     []student_events.StudentEventsData          `json:"student_events"`
}

func (s *ScheduleSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	student, err := middlewares.GetStudentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	studentClassID, err := middlewares.GetStudentClassIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	q := r.URL.Query().Get("time")
	if q == "" {
		http.Error(w, "'time' query params needed", http.StatusUnauthorized)
		return
	}
	t, err := time.Parse(utils.LayoutFormat, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		s.StudentEvents, err = student_events.StudentEventDB.GetByStudentIDAndTimestamp(student.ID, t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}()
	go func() {
		defer wg.Done()
		s.SchoolClassEvents, err = school_class_events.SchoolClassEventDB.GetBySchoolClassIDAndTimestamp(studentClassID, t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}()
	wg.Wait()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/siswa/schedule", middlewares.StudentMiddleware(
		middlewares.StudentClassMiddlewares(
			(&ScheduleSiswa{}).ServeHTTP),
	),
	)
}
