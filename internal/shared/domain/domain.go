package domain

import "github.com/google/uuid"

// Map represents a map that represents the format of the http response
type Map map[string]any

// UuID is a type of string representing a identifier value object
type UuID string

// EnsureIdIsValid validate if the format the values is a UuID
func (i UuID) EnsureIdIsValid() error {
	_, err := uuid.Parse(string(i))
	return err
}

// GenerateUuID generate a new UuID.
func (UuID) GenerateUuID() string {
	return uuid.New().String()
}
