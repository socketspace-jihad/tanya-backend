package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
)

func GetStudentFromRequestContext(r *http.Request) (*student_profiles.StudentProfilesData, error) {
	userRole, ok := r.Context().Value(ContextKey("student_profiles")).(student_profiles.StudentProfilesData)
	if !ok {
		return nil, errors.New("jwt token not valid")
	}
	return &userRole, nil
}

func StudentMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		t, err := tokenParser(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if err := token.Claims.Valid(); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		data, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, errors.New("invalid jwt token").Error(), http.StatusUnauthorized)
			return
		}
		var u student_profiles.StudentProfilesData
		if err := mapstructure.Decode(data, &u); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextKey("student_profiles"), u)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
