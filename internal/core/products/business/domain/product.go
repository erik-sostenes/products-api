package domain

import (
	"net/url"

	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
)

// Sales represents an Entity that composes the Product Entity
type Sales struct {
	Days   int
	Amount float64
}

// Product represents an Entity of my domain
type Product struct {
	Id        string
	Title     string
	ImageURL  string
	Price     float64
	Rating    float64
	Offer     bool
	Available bool
	Sales     Sales
}

// NewProduct returns a product instance, only if the values of the object are correct
func NewProduct(id, title, imageUrl string, price, rating float64, offer, available bool, days int, amount float64) (Product, error) {
	if err := URL(imageUrl).ensureUrlIsValid(); err != nil {
		return Product{}, wrongs.StatusBadRequest(err.Error())
	}

	return Product{
		Id:        id,
		Title:     title,
		Price:     price,
		ImageURL:  imageUrl,
		Offer:     offer,
		Available: available,
		Rating:    rating,
		Sales: Sales{
			Days:   days,
			Amount: amount,
		},
	}, nil
}

// URL is a type of string representing a URL value object
type URL string

// ensureUrlIsValid method that validates the format of the url
func (u URL) ensureUrlIsValid() error {
	_, err := url.ParseRequestURI(string(u))

	return err
}
