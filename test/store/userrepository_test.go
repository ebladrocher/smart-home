package store_test

import (
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/ebladrocher/smarthome/system/store/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	thisDB, teardown := db.TestDB(t, dataBaseURL)
	defer teardown("users")

	s := db.NewStore(thisDB)
	u := models.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	thisDB, teardown := db.TestDB(t, dataBaseURL)
	defer teardown("users")

	s := db.NewStore(thisDB)
	email := "user@user.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)


}