package db

import (
	"database/sql"
	"github.com/ebladrocher/smarthome/models"
	"github.com/ebladrocher/smarthome/system/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {
	if err := u.CheckPassword(); err != nil {
		return err
	}

	return  r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) ( *models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
		).Scan(
			&u.ID,
			&u.Email,
			&u.EncryptedPassword,
		); err != nil {
				if err ==  sql.ErrNoRows {
					return nil, store.ErrRecordNotFound
				}

				return nil, err
		}

	return u,nil
}

