package teststore

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
	"sync"
)

// UserLocalStorage ...
type UserLocalStorage struct {
	users map[string]*models.User
	mutex *sync.Mutex
}

// NewUserLocalStorage ...
func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*models.User),
		mutex: new(sync.Mutex),
	}
}

// CreateUser ...
func (s *UserLocalStorage) CreateUser(user *models.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

// GetUser ...
func (s *UserLocalStorage) GetUser(email, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, store.ErrUserNotFound
}
