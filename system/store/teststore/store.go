package teststore

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// NewStore ...
func NewStore() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository !=  nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store:s,
		users:make(map[string]*models.User),
	}

	return s.userRepository
}
