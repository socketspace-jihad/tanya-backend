package user_topics

import "github.com/socketspace-jihad/tanya-backend/models/user"

type UserTopicsData struct {
	ID            uint `json:"id"`
	user.UserData `json:"user"`
	Topic         string `json:"topic"`
}
