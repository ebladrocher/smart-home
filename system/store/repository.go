package store

import "github.com/ebladrocher/smarthome/models"

// UserRepository ...
type UserRepository interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
}
