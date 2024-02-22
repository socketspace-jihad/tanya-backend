package tugas_siswa

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/student_tasks"
)

type TugasSiswa struct{}

func (t *TugasSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	student, err := middlewares.GetStudentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tugas, err := student_tasks.StudentTasksDB.GetByStudentID(student.ID)
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
