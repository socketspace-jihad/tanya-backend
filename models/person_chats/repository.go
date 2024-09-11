package person_chats

type PersonChatsRepository interface {
	Save(*PersonChatsData) error
	GetBySourceUserID(uint) ([]PersonChatsData, error)
}
