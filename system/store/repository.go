package store

import "github.com/ebladrocher/smarthome/models"

// UserRepository ...
type UserRepository interface {
	Create(*models.User) error
	FindByEmail(string) (*models.User, error)
}
