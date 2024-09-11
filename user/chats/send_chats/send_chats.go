package sendchats

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/global_chats_detail"
)

type SendChats struct{}

func (s *SendChats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var chat global_chats_detail.GlobalChatsDetailData
	if err := json.Unmarshal(body, &chat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chat.SourceUser = *user
	if err := global_chats_detail.GlobalChatDetailsDB.Save(&chat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chat)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/chats/send",
		auth.AuthMiddlewareHandler(&SendChats{}),
	)
}
