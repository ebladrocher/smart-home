package store

import "github.com/ebladrocher/smarthome/models"

type UseCase interface {
	// User
	SignUp(email, password string) error
	SignIn(email, password string) (*models.User, error)
}
