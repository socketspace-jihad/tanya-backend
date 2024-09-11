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
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type RegistrasiGuru struct{}

func (rs *RegistrasiGuru) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	seedTeacherProfile := teacher_profiles.TeacherProfilesData{}
	if err := json.Unmarshal(body, &seedTeacherProfile); err != nil {
		log.Println("ERROR UNMARSHAL", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRole := &user_roles.UserRolesData{
		UserData: user.UserData{
			ID: u.ID,
		},
		RolesData: roles.RolesData{
			ID: roles.GuruRolesID,
		},
	}
	if err := user_roles.UserRolesDB.Save(userRole); err != nil {
		log.Println("ERROR SAVE USER ROLE", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seedTeacherProfile.UserRolesData = *userRole
	if err := teacher_profiles.TeacherProfilesDB.Save(&seedTeacherProfile); err != nil {
		log.Println("ERROR SAVE SEED STUDENT PROFILES", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRole.CurrentTeacherProfilesID = seedTeacherProfile.ID
	if err := user_roles.UserRolesDB.Update(userRole); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(seedTeacherProfile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/teacher/profile/register", auth.AuthMiddlewareHandler(&RegistrasiGuru{}))
}
