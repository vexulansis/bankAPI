package storage

import (
	"bankAPI/internal/model"
	"database/sql"
)

// Get user info from sql.Rows
func ParseUserInfo(rows *sql.Rows) ([]*model.User, error) {
	users := []*model.User{}
	for rows.Next() {
		u := new(model.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Get auth info from sql.Rows
func ParseAuthInfo(rows *sql.Rows) ([]*model.Auth, error) {
	auths := []*model.Auth{}
	for rows.Next() {
		a := new(model.Auth)
		err := rows.Scan(&a.User_id, &a.Login, &a.Password)
		if err != nil {
			return nil, err
		}
		auths = append(auths, a)
	}
	return auths, nil
}
