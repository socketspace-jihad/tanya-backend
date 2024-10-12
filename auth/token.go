package auth

import (
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
)

type Token struct{}

func (tkn *Token) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}

func init() {
	http.DefaultServeMux.Handle("/v1/account/refresh-token", auth.AuthMiddlewareHandler(&Token{}))
}
