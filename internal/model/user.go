package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Auth  Auth   `json:"auth"`
}
type Auth struct {
	User_id           int    `json:"user_id"`
	Login             string `json:"login"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encryptedpassword"`
}

// Preparing user before creation
func (u *User) BeforeCreate() error {
	if len(u.Auth.Password) > 0 {
		enc, err := HashPassword(u.Auth.Password)
		if err != nil {
			return err
		}
		u.Auth.EncryptedPassword = enc
		u.Sanitize()
	}
	return nil
}

// Encrypting password method
func HashPassword(pas string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Cleansing user password field
func (u *User) Sanitize() {
	u.Auth.Password = ""
}
