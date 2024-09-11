package global_chats

type GlobalChatsRepository interface {
	Save(*GlobalChatsData) error
	GetBySourceUserID(uint) ([]GlobalChatsData, error)
}
