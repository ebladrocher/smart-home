package store_test

import (
	"github.com/stretchr/testify/assert"
	"smarthome/model"
	"smarthome/system/store"
	"testing"
)

func TestUserRepository(t *testing.T)  {
	s, teardown := store.TestStore(t)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@user.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T)  {
	s, teardown := store.TestStore(t)
	defer teardown("users")

	email := "user@user.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@user.org",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)


}