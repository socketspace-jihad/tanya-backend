package kegiatan_siswa_parent

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type KegiatanSiswaParent struct{}

func (k *KegiatanSiswaParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		studentID := r.URL.Query().Get("student_profiles_id")
		if studentID == "" {
			http.Error(w, "student_profiles_id must be defined on query params", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(studentID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		student, err := student_profiles.StudentProfilesDB.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		classEvents, err := school_class_events.SchoolClassEventDB.GetBySchoolClassID(student.CurrentSchoolData.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(classEvents)
	}
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/student/events", middlewares.ParentMiddleware(
		&KegiatanSiswaParent{}),
	)
}
