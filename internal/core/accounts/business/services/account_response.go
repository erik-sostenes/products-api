package services

import "github.com/erik-sostenes/products-api/internal/shared/domain"

// AccountResponse represent an DTO(Data Transfer Object)
type AccountResponse struct {
	AccountId       string
	AccountUserName string
	AccountPassword string
	AccountDetails  domain.Map
}
