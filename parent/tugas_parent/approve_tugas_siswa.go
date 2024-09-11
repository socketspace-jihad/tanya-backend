package tugas_parent

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/class_tasks_student_parent"
)

type ApproveTugasSiswa struct{}

func (t *ApproveTugasSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var tasks class_tasks_student_parent.ClassTasksStudentParentData
	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		log.Println("decoder error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks.ParentProfilesData = *parent
	if err := class_tasks_student_parent.ClassTasksStudentParentDB.Save(&tasks); err != nil {
		log.Println("saving error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/student/tugas/approve", middlewares.ParentMiddleware((&ApproveTugasSiswa{})))
}
