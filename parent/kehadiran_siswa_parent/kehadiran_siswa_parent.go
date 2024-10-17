package kehadiran_siswa_parent

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
)

type KehadiranSiswaParent struct{}

func (k *KehadiranSiswaParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		http.Error(w, "status unauthorized", http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query().Get("student_profiles_id")
		if params == "" {
			http.Error(w, "student_profiles_id must be defined as a parameter", http.StatusBadRequest)
			return
		}
		studentID, err := strconv.Atoi(params)
		if err != nil {
			http.Error(w, "student_profiles_id not valid", http.StatusBadRequest)
			return
		}
		presensi, err := student_presensi.StudentPresensiDB.GetByStudentProfilesID(uint(studentID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(presensi)
	default:
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/parent/student/presensi", middlewares.ParentMiddleware(&KehadiranSiswaParent{}))
}
