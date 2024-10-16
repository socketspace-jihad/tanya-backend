package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type RolesLogin struct {
	Token string `json:"token"`
}

func (a *RolesLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := auth.GetUser(r)
	if err != nil {
		http.Error(w, errors.New("jwt token invalid").Error(), http.StatusUnauthorized)
		return
	}
	var userRole user_roles.UserRolesData

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.Unmarshal(body, &userRole); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRole, err = user_roles.UserRolesDB.GetByRoleAndUserID(userRole.RolesData.ID, u.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userRole.ID == 0 {
		http.Error(w, errors.New("user_roles not found within this user").Error(), http.StatusNotFound)
		return
	}

	switch userRole.RolesData.ID {
	// guru
	case 1:
		teacher, err := teacher_profiles.TeacherProfilesDB.GetByUserRoleID(userRole.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		teacher.UserRolesData = userRole
		teacher.UserRolesData.UserData = *u
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, teacher)
		signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		a.Token = signedToken
		json.NewEncoder(w).Encode(a)

	// siswa
	case 2:
		student, err := student_profiles.StudentProfilesDB.GetByUserRoleID(userRole.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		student.UserRolesData = userRole
		student.UserRolesData.UserData = *u
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, student)
		signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		a.Token = signedToken
		json.NewEncoder(w).Encode(a)

	// orang tua
	case 3:
		parent, err := parent_profiles.ParentProfilesDB.GetByUserRoleID(userRole.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		parent.UserRolesData = userRole
		parent.UserRolesData.UserData = *u
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, parent)
		signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		a.Token = signedToken
		json.NewEncoder(w).Encode(a)
	// admin
	case 4:
	}

}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/account/roles/login", auth.AuthMiddlewareHandler(&RolesLogin{}))
}
