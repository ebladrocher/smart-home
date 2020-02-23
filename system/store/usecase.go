package store

type UseCase interface {
	// User
	SignUp(username, password string) error
}
