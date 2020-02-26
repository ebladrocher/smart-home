package config

import "time"

// Token ...
type Token struct {
	HashSalt   string        `json:"hash_salt"`
	SigningKey string        `json:"signing_key"`
	Token      time.Duration `json:"token_ttl"`
}
