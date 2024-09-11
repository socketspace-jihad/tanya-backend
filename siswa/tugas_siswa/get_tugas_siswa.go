package tugas_siswa

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student"
)

type TugasSiswa struct{}

func (t *TugasSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	student, err := middlewares.GetStudentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tugas, err := class_tasks_student.ClassTasksStudentDB.GetByStudentProfilesID(student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tugas)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/siswa/tugas", middlewares.StudentMiddleware((&TugasSiswa{}).ServeHTTP))
}
