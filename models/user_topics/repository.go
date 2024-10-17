package user_topics

type UserTopicsRepository interface {
	Save(*UserTopicsData) error
	GetByID(uint) (*UserTopicsData, error)
	GetByUserID(uint) ([]UserTopicsData, error)
	GetByTopicName(string) ([]UserTopicsData, error)
}
