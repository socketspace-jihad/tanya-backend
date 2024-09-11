package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"
	"github.com/socketspace-jihad/tanya-backend/models/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthLogin struct {
	Token string `json:"token"`
}

func (a *AuthLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u user.UserData

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.Unmarshal(body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := user.UserDB.Repo.GetByEmail(u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(u.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &data)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.Token = signedToken
	json.NewEncoder(w).Encode(a)
}

func init() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	http.DefaultServeMux.Handle("/v1/account/login", c.Handler(&AuthLogin{}))
}
