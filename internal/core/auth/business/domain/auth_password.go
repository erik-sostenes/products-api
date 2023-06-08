package domain

import (
	"golang.org/x/crypto/bcrypt"
)

// AuthPassword (Value Object) represents the auth password
type AuthPassword struct {
	hashedPassword string
}

// NewAuthPassword returns an instance of AuthPassword
func NewAuthPassword(value string) (AuthPassword, error) {
	return AuthPassword{
		value,
	}, nil
}

// Equals compares the hashed password with which provided the user,
// if the two passwords match returns nil, but if they do not match an error return
func (a AuthPassword) Equals(password AuthPassword) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.hashedPassword), []byte(password.String()))
}

func (a AuthPassword) String() string {
	return a.hashedPassword
}
