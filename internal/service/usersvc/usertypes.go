package usersvc

import db "github.com/axilock/axilock-backend/internal/db/sqlc"

type CreateUserRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateUserResponse struct {
	User *db.User
	Init bool
}
