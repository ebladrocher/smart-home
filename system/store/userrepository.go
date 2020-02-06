package store

import "smarthome/models"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) ( *models.User, error) {

	if err := u.CheckPassword(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
		).Scan(&u.ID); err !=nil {
			return nil,err
	}

	return u,nil
}

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
				return nil, err
	}

	return u,nil
}

