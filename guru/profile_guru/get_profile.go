package profile_siswa

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
)

type ProfileGuru struct{}

func (p *ProfileGuru) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	profile, err := teacher_profiles.TeacherProfilesDB.GetByID(g.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/guru/profile", middlewares.TeacherMiddleware(&ProfileGuru{}))
}
