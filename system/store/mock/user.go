package mock

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(email, password string) (*models.User, error) {
	args := s.Called(email, password)

	return args.Get(0).(*models.User), args.Error(1)
}
