package get_chats

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/person_chats"
)

type GetChats struct{}

func (g *GetChats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	chats, err := person_chats.PersonChatsDB.GetBySourceUserID(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chats)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/chats",
		auth.AuthMiddlewareHandler(&GetChats{}),
	)
}
