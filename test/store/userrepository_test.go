package store_test

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store/postgrestore"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestUserRepository_Create(t *testing.T) {
	thisDB, teardown := postgrestore.TestDB(t, dataBaseURL)
	defer teardown("users")

	s := postgrestore.NewUserRepository(thisDB)
	u := &models.User{
		Email: "testuser",
		Password:"testpass",
	}
	assert.NoError(t, s.CreateUser(u))
	assert.NotNil(t, u)
}

//func TestUserRepository_FindByEmail(t *testing.T) {
//	thisDB, teardown := postgrestore.TestDB(t, dataBaseURL)
//	defer teardown("users")
//
//	s := postgrestore.NewStore(thisDB)
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