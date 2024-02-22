package profile_siswa

import (
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
)

type ProfileSiswa struct{}

func (p *ProfileSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := auth.GetUserRoles(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Write([]byte("GET PROFILE SISWA"))
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/siswa/profile", auth.RolesMiddlewareHandler(&ProfileSiswa{}))
}
