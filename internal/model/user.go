package model

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Auth  Auth   `json:"auth"`
}
type Auth struct {
	User_id  int    `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
