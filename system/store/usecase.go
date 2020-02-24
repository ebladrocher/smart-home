package store

type UseCase interface {
	// User
	SignUp(email, password string) error
}
