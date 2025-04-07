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

type HasAccessUserRequest struct {
	Token string `json:"token"`
}

type HasAccessUserResponse struct {
	UserId    uint `json:"user_id"`
	HasAccess bool `json:"has_access"`
}
