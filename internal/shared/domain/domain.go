package domain

import (
	"fmt"

	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/google/uuid"
)

// Map represents a map that represents the format of the http response
type Map map[string]any

// UuID is a type of string representing a identifier value object
type UuID string

// EnsureIdIsValid validate if the format the values is a UuID
func (i UuID) EnsureIdIsValid() error {
	_, err := uuid.Parse(string(i))
	if err != nil {
		return wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", string(i)))
	}
	return nil
}

// GenerateUuID generate a new UuID.
func (UuID) GenerateUuID() string {
	return uuid.New().String()
}
