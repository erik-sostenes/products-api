package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/ports"
	"github.com/erik-sostenes/products-api/internal/shared/domain/wrongs"
	"github.com/go-sql-driver/mysql"
)

// ProductStorer implements the ports.ProductStorer interface and persist the data in mysql
type ProductStorer struct {
	DB *sql.DB
}

// NewProductStorer returns an instance ports.ProductStorer
func NewProductStorer(DB *sql.DB) ports.ProductStorer {
	return ProductStorer{
		DB: DB,
	}
}

// Save method that persist the product a mysql
func (p ProductStorer) Save(ctx context.Context, identifier string, product domain.Product) (err error) {
	var sqlInsertProduct = "insert into products(id, title, image_url, price, rating, offer, available, sales_days, sales_amount) values(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = p.DB.Exec(sqlInsertProduct,
		identifier,
		product.Title,
		product.ImageURL,
		product.Price,
		product.Rating,
		product.Offer,
		product.Available,
		product.Sales.Days,
		product.Sales.Amount,
	)

	if code, ok := err.(*mysql.MySQLError); ok {
		//NOTE: Error Code: 1062. Duplicate entry key
		if code.Number == 1062 {
			return wrongs.StatusBadRequest(fmt.Sprintf("Resource with id %v already existing.", identifier))
		}

		err = errors.New("An error has occurred while adding a new product record.")
	}
	return err
}

// Find method searches for all the records in mysql and returns them in a slice
func (p ProductStorer) Find(ctx context.Context) ([]domain.Product, error) {
	var sqlSelectProducts = "select * from products"
	res, err := p.DB.Query(sqlSelectProducts)
	defer res.Close()

	if err != nil {
		return nil, errors.New("An error occurred while obtaining the products.")
	}

	var products []domain.Product
	for res.Next() {
		var product domain.Product

		if err := res.Scan(&product); err != nil {
			return nil, errors.New("An error occurred while obtaining the products.")
		}
		products = append(products, product)
	}
	return products, nil
}

// FindById method searches for a record by identifier in mysql
func (p ProductStorer) FindById(ctx context.Context, identifier string) (domain.Product, error) {
	var product domain.Product

	var sqlSelectProduct = "select * from products p where p.id = ?"
	err := p.DB.QueryRowContext(ctx, sqlSelectProduct, identifier).Scan(&product)
	if err != nil {
		return domain.Product{}, errors.New("An error occurred while obtaining the product.")
	}

	if product == (domain.Product{}) {
		return domain.Product{}, wrongs.StatusNotFound(fmt.Sprintf("Resource with identifier %v does no exist.", identifier))
	}

	return product, nil
}

// Delete method deletes a record by identifier in mysql
func (p ProductStorer) Delete(ctx context.Context, identifier string) error {
	var sqlDeleteProduct = "delete from products p where p.id = ?"

	row, err := p.DB.Exec(sqlDeleteProduct, identifier)
	if err != nil {
		return errors.New("An error occurred while deleting the products.")
	}

	if rowAffected, _ := row.RowsAffected(); rowAffected != 1 {
		return wrongs.StatusNotFound(fmt.Sprintf("Resource with identifier %v does no exist.", identifier))
	}

	return nil
}
