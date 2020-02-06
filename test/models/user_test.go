package models_test

import (
	"github.com/stretchr/testify/assert"
	"smarthome/models"
	"testing"
)

func TestUser_CheckPassword(t *testing.T)  {
	u := models.TestUser(t)
	assert.NoError(t, u.CheckPassword())
	assert.NotEmpty(t, u.EncryptedPassword)
}