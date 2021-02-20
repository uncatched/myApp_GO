package models

// AuthUser ...
type AuthUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
