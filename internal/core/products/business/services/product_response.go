package services

import "github.com/erik-sostenes/products-api/internal/core/products/business/domain"

// ProductResponse represent an DTO(Data Transfer Object)
type ProductResponse struct {
	Id        string
	Title     string
	ImageURL  string
	Price     float64
	Rating    float64
	Offer     bool
	Available bool
	Sales     domain.Sales
}
