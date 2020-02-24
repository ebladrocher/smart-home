package teststore__test

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/ebladrocher/smarthome/system/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := teststore.NewUserLocalStorage()

	id1 := "id"

	user := &models.User{
		ID:       id1,
		Email: "user",
		Password: "password",
	}

	err := s.CreateUser(user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser("user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.GetUser("user", "")
	assert.Error(t, err)
	assert.Equal(t, err, store.ErrUserNotFound)
}

//func TestUserRepository_FindByEmail(t *testing.T)  {
//	s := teststore.NewStore()
//	email := "user@user.org"
//	_, err := s.User().FindByEmail(email)
//	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
//
//	u := models.TestUser(t)
//	u.Email = email
//	s.User().Create(u)
//	u, err = s.User().FindByEmail(email)
//	assert.NoError(t, err)
//	assert.NotNil(t, u)
//
//
//}


