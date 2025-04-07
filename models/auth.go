package models

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
