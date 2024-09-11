package person_chats

import (
	"github.com/socketspace-jihad/tanya-backend/models/global_chats"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type PersonChatsData struct {
	ID                           uint          `json:"id"`
	SourceUser                   user.UserData `json:"source_user,omitempty"`
	TargetUser                   user.UserData `json:"target_user,omitempty"`
	global_chats.GlobalChatsData `json:"global_chats"`
}
