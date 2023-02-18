package storage

import "bankAPI/internal/model"

// Creating user with transaction
func (storage *Storage) CreateUser(u *model.User) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx.QueryRow("insert into users default values returning id").Scan(&u.ID)
	_, err = tx.Exec("insert into auth (user_id, login, password) values ($1, $2, $3)", u.ID, u.Auth.Login, u.Auth.EncryptedPassword)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
