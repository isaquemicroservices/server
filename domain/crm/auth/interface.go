package auth

// IAuth interface of methods to auth
type IAuth interface {
	CreateUser(user *User) (err error)
	Login(user *Credentials) (*User, error)
}
