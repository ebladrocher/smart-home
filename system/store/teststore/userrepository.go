package teststore

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*models.User
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {
	if err := u.CheckPassword(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u, ok:= r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
