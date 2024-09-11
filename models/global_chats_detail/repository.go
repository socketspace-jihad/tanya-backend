package global_chats_detail

type GlobalChatsDetailRepository interface {
	Save(*GlobalChatsDetailData) error
	GetByGlobalChatsID(uint) ([]GlobalChatsDetailData, error)
}
