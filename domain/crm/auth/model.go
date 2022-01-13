package auth

import "time"

// User model of user
type User struct {
	Id        int64
	Name      string
	Email     string
	Passw     string
	Token     string
	CreateAt  time.Time
	UpdatedAt time.Time
}

// Credentials credentials for authentication
type Credentials struct {
	Email string `json:"email"`
	Passw string `json:"passw"`
}
