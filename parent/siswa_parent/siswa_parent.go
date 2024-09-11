package siswa_parent

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/parent_student"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type SiswaParent struct{}

func (s *SiswaParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	case http.MethodPost:
		parent, err := middlewares.GetParentFromRequestContext(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		student := student_profiles.StudentProfilesData{}
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		parentStudent := parent_student.ParentStudentData{
			StudentProfilesData: student,
			ParentProfilesData:  *parent,
		}
		if err := parent_student.ParentStudentDB.Save(&parentStudent); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parentStudent)
	}
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/student", middlewares.ParentMiddleware(
		&SiswaParent{}),
	)
}
