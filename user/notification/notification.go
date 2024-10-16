package notification

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
)

type Notification struct{}

func (n *Notification) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		notifications, err := notification.NotificationDB.GetByUserID(
			user.ID,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notifications)
	}
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/notification",
		auth.AuthMiddlewareHandler((&Notification{})),
	)
}
