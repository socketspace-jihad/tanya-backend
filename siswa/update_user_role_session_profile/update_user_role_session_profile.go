package update_user_role_session_profile

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/user_roles"
)

type UpdateUserRoleSessionProfile struct{}

func (p *UpdateUserRoleSessionProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userRoles user_roles.UserRolesData

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &userRoles); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := user_roles.UserRolesDB.Update(&userRoles); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/user/roles/student/profiles/update", auth.AuthMiddlewareHandler(&UpdateUserRoleSessionProfile{}))
}
