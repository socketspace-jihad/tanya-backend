package user_token

import "github.com/socketspace-jihad/tanya-backend/models/user"

type UserTokenData struct {
	ID            uint `json:"id"`
	user.UserData `json:"user"`
	Token         string `json:"token"`
	CreatedAt     string `json:"created_at"`
	Status        uint8  `json:"status"`
}
