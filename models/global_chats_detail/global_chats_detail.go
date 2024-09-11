package global_chats_detail

import (
	"time"

	"github.com/socketspace-jihad/tanya-backend/models/global_chats"
	"github.com/socketspace-jihad/tanya-backend/models/user"
)

type GlobalChatsDetailData struct {
	ID                           uint          `json:"id"`
	SourceUser                   user.UserData `json:"source_user"`
	TargetUser                   user.UserData `json:"target_user,omitempty"`
	global_chats.GlobalChatsData `json:"global_chats,omitempty"`
	CreatedAt                    time.Time `json:"created_at"`
	Contents                     string    `json:"contents"`
	Attachment                   string    `json:"attachment"`
}
