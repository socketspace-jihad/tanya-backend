package tugas_parent

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student"
)

type TugasParent struct{}

func (t *TugasParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tasks, err := class_tasks_student.ClassTasksStudentDB.GetByParentProfilesID(parent.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/student/tugas", middlewares.ParentMiddleware((&TugasParent{})))
}
