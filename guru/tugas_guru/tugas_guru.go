package tugas_guru

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks"
)

type TugasGuru struct{}

func (t *TugasGuru) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		teacher, err := middlewares.GetTeacherFromRequestContext(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tugas, err := class_tasks.ClassTasksDB.GetByTeacherProfilesID(teacher.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&tugas)
	case http.MethodPost:
		return
	}
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/guru/tugas", middlewares.TeacherMiddleware(
		&TugasGuru{}),
	)
}
