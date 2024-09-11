package registrasi_siswa

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/roles"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type RegistrasiSiswa struct{}

func (rs *RegistrasiSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middlewares.ContextKey("user")).(user.UserData)
	if !ok {
		http.Error(w, errors.New("jwt token invalid").Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seedStudentProfile := student_profiles.StudentProfilesData{}
	if err := json.Unmarshal(body, &seedStudentProfile); err != nil {
		log.Println("ERROR UNMARSHAL", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRole := &user_roles.UserRolesData{
		UserData: user.UserData{
			ID: u.ID,
		},
		RolesData: roles.RolesData{
			ID: roles.SiswaRolesID,
		},
	}
	if err := user_roles.UserRolesDB.Save(userRole); err != nil {
		log.Println("ERROR SAVE USER ROLE", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// create seed student profile
	seedStudentProfile.UserRolesData = *userRole
	if err := student_profiles.StudentProfilesDB.Save(&seedStudentProfile); err != nil {
		log.Println("ERROR SAVE SEED STUDENT PROFILES", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRole.CurrentStudentProfilesID = seedStudentProfile.ID
	if err := user_roles.UserRolesDB.Update(userRole); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(seedStudentProfile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/student/profile/register", auth.AuthMiddlewareHandler(&RegistrasiSiswa{}))
}
