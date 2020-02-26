package models

// User ...
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// Validate ...
//func Valedate() error {
//
//}

//// Encrypt string
//func encryptString(s string) (string, error) {
//	e, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
//	if err != nil {
//		return "", err
//	}
//	return string(e), nil
//}
