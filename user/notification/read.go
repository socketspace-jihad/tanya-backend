package notification

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
)

type NotificationRead struct{}

func (n *NotificationRead) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var notif notification.NotificationData
		if err := json.Unmarshal(body, &notif); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		notif.UserData = *user
		if err := notification.NotificationDB.UpdateRead(&notif); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode([]byte("ok"))
	}
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/notification/read",
		auth.AuthMiddlewareHandler((&NotificationRead{})),
	)
}
