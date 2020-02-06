package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID               int
	Email            string
	Password         string
	EncryptedPassword string
}

//func (u  *User) Validate() error {
//	return nil
//}

// Check Password ....
func (u *User) CheckPassword()  error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// Encrypt string
func encryptString(s string) (string, error) {
	e, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(e), nil
}