package auth

import "time"

// User struct for user
type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Passw     string    `json:"passw,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CredentialsReq credentials for authentication
type CredentialsReq struct {
	Email string `json:"email"`
	Passw string `json:"passw"`
}
