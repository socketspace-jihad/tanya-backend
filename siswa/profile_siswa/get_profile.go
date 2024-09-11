package profile_siswa

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

type ProfileSiswa struct{}

func (p *ProfileSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	studentProfiles, err := auth.GetUserRoles(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	profile, err := student_profiles.StudentProfilesDB.GetByID(studentProfiles.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/siswa/profile", auth.RolesMiddlewareHandler(&ProfileSiswa{}))
}
