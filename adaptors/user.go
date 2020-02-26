package adaptors

import (
	"crypto/sha1"
	"fmt"
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

// AuthUseCase ...
type AuthUseCase struct {
	userRepo       store.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// NewAuthUseCase ...
func NewAuthUseCase(
	userRepo store.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

// SignUp ...
func (a *AuthUseCase) SignUp(email, password string) error {
	if err := Validate(email,password); err != nil {
		return err
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Email: email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(user)
}

// SignIn ...
func (a *AuthUseCase) SignIn(email, password string) (*models.User, error) {
	if err := Validate(email,password); err != nil {
		return nil, err
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	u, err := a.userRepo.GetUser(email)
	if err != nil || (u.Password != password){
		return nil, store.ErrUserNotFound
	}

	return u, nil
}

// Validate ...
func Validate(e, p string) error {

	if err := validation.Validate(e, validation.Required, is.Email); err != nil {
		return err
	}

	if err := validation.Validate(p, validation.Required, validation.Length(6,100)); err != nil {
		return err
	}
	return nil
}
