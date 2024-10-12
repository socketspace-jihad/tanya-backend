package student_class

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
)

type StudentClass struct{}

func (e *StudentClass) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("school_class_id")
	if data == "" {
		return
	}
	classID, err := strconv.Atoi(data)
	if err != nil {
		return
	}
	students, err := student_profiles.StudentProfilesDB.GetBySchoolClassID(uint(classID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func init() {
	http.DefaultServeMux.Handle("/v1/class/student", auth.AuthMiddlewareHandler(&StudentClass{}))
}
