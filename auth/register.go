package auth

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type Register struct{}

func (a *Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u user.UserData
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.UserDB.Repo.Save(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func init() {
	http.DefaultServeMux.Handle("/v1/account/register", &Register{})
}
