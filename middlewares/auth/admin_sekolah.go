package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/roles"
)

func AdminSekolahMiddleware(h http.Handler) http.HandlerFunc {
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
		userRole, err := GetUserRoles(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if userRole.RoleID != roles.AdminSekolahRolesID {
			http.Error(w, "the role is not school admin", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), middlewares.ContextKey("user_roles"), token.Claims)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
