package adaptors

import "github.com/stretchr/testify/mock"

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(email, password string) error {
	args := m.Called(email, password)

	return args.Error(0)
}
