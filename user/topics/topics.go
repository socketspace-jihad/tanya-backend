package topics

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares/auth"
	"github.com/socketspace-jihad/tanya-backend/models/user_topics"
)

type UserTopics struct{}

func (u *UserTopics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := auth.GetUser(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		topics, err := user_topics.UserTopicsDB.GetByUserID(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(topics)
		return
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var userTopic user_topics.UserTopicsData
		if err := json.Unmarshal(body, &userTopic); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userTopic.UserData = *user
		user_topics.UserTopicsDB.Save(&userTopic)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/user/topics",
		auth.AuthMiddlewareHandler(&UserTopics{}),
	)
}
