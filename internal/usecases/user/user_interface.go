package user_usecase

type UserUsecaseI interface {
	FirstAttempt(inp string) (int, string)
	LoginWithEmail(email, password string) error
	LoginWithPhone(phone string) error
	SignupWithEmail(email, password string) error
	SignupWithPhone(Phone, code string) error
	SaveToken(id, token string) error
	TokenExists(id, token string) bool
}
