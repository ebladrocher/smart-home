package store_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T)  {
	s, teardown := store.TestStore(t)
	defer teardown("users")

	u, err := s.User().Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T)  {
	s, teardown := store.TestStore(t)
	defer teardown("users")

	email := "user@user.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)


}