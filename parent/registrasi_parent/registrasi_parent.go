package registrasi_parent

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/roles"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type RegistrasiParent struct{}

func (rs *RegistrasiParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := auth.GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seedParentProfile := parent_profiles.ParentProfilesData{}
	if err := json.Unmarshal(body, &seedParentProfile); err != nil {
		log.Println("ERROR UNMARSHAL", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRole := &user_roles.UserRolesData{
		UserData: user.UserData{
			ID: u.ID,
		},
		RolesData: roles.RolesData{
			ID: roles.OrangTuaRolesID,
		},
	}
	if err := user_roles.UserRolesDB.Save(userRole); err != nil {
		log.Println("ERROR SAVE USER ROLE", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seedParentProfile.UserRolesData = *userRole
	if err := parent_profiles.ParentProfilesDB.Save(&seedParentProfile); err != nil {
		log.Println("ERROR SAVE SEED PARENT PROFILES", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(seedParentProfile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/profile/register", auth.AuthMiddlewareHandler(&RegistrasiParent{}))
}
