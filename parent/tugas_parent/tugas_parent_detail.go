package tugas_parent

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student"
)

type TugasParentDetail struct{}

func (t *TugasParentDetail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := r.URL.Query().Get("class_tasks_student_id")
	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks, err := class_tasks_student.ClassTasksStudentDB.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/student/tugas/detail", middlewares.ParentMiddleware((&TugasParentDetail{})))
}
