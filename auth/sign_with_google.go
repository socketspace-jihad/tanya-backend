package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type SignWithGoogle struct {
	Token string `json:"token"`
}

func (s *SignWithGoogle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u user.UserData
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("ERROR DECODE")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := user.UserDB.Repo.GetByEmail(u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// do register
	if data.ID == 0 {
		if err := user.UserDB.Repo.Save(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.ID = u.ID
	}
	data.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &data)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.Token = signedToken
	json.NewEncoder(w).Encode(s)
	return
}

func init() {
	http.DefaultServeMux.Handle("/v1/account/sign/google", &SignWithGoogle{})
}
