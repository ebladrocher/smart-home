package postgrestore

import (
	"database/sql"
	"github.com/ebladrocher/smarthome/models"
	"strconv"
)

type User struct {
	ID       int
	Email    string
	Password string
}

// UserRepository ...
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository ...
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create ...
func (r UserRepository) CreateUser(user *models.User) error {
	model := toPostgresUser(user)

	return  r.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		model.Email,
		model.Password,
	).Scan(&model.ID)
}

// FindByEmail ...
func (r UserRepository) GetUser(email string) (*models.User, error) {
	user := new(User)

	if err := r.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	); err != nil {
		if err ==  sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return toModel(user),nil
}

// toPostgresUser ...
func toPostgresUser(u *models.User) *User {
	return &User{
		Email: u.Email,
		Password: u.Password,
	}
}

// toModel ...
func toModel(u *User) *models.User {
	return &models.User{
		ID:       strconv.Itoa(u.ID),
		Email: u.Email,
		Password: u.Password,
	}
}
