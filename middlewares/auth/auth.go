package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

func tokenParser(token string) (string, error) {
	if token == "" {
		return "", errors.New("authorization header cannot be nil")
	}
	parsed := strings.Split(token, " ")
	if len(parsed) < 2 {
		return "", errors.New("invalid token format")
	}
	return parsed[1], nil
}

func GetUser(r *http.Request) (*user.UserData, error) {
	user, ok := r.Context().Value(middlewares.ContextKey("user")).(user.UserData)
	if !ok {
		return nil, errors.New("jwt token not valid")
	}
	return &user, nil
}

func AuthMiddlewareHandler(h http.Handler) http.HandlerFunc {
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
		var u user.UserData
		if err := mapstructure.Decode(data, &u); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), middlewares.ContextKey("user"), u)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
