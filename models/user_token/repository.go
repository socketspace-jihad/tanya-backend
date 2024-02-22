package user_token

type UserTokenRepository interface {
	ValidateConnection() error
	Save(UserTokenData) error
	List() []UserTokenData
	GetByID(uint) (UserTokenData, error)
	GetByToken(string) (UserTokenData, error)
}
