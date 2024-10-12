package get_detail_chats

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/global_chats_detail"
)

type GetDetailChats struct{}

func (g *GetDetailChats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query().Get("global_chats_id")
	globalChatsID, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chats, err := global_chats_detail.GlobalChatDetailsDB.GetByGlobalChatsID(uint(globalChatsID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chats)
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/chats/detail",
		auth.AuthMiddlewareHandler(&GetDetailChats{}),
	)
}
