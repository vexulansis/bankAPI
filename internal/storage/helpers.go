package storage

import (
	"bankAPI/internal/model"
)

// Verify user existence
func (storage *Storage) CheckUser(u *model.User) (bool, error) {
	rows, err := storage.Search(
		WithData(u),
		WithTable("auth"),
		WithReferenceName("login"),
		WithReference(u.Auth.Login),
	)
	if err != nil {
		return false, err
	}
	auths, err := ParseAuthInfo(rows)
	if err != nil {
		return false, err
	}
	if len(auths) == 0 {
		return false, nil
	}
	return true, nil
}

// Verify user password
func (storage *Storage) VerifyPassword(u *model.User) (bool, error) {
	rows, err := storage.Search(
		WithData(u),
		WithTable("auth"),
		WithReferenceName("login"),
		WithReference(u.Auth.Login),
	)
	if err != nil {
		return false, err
	}
	auths, err := ParseAuthInfo(rows)
	if err != nil {
		return false, err
	}
	ref := auths[0]
	if ref.Password == u.Auth.Password {
		return true, nil
	}
	return false, nil
}
