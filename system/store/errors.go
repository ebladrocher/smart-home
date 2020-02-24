package store

import "errors"

var (
	// ErrorRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
	// ErrUserNotFound  ...
	ErrUserNotFound = errors.New("user not found")
	// ErrInvalidAccessToken ...
	ErrInvalidAccessToken = errors.New("invalid access token")
)
