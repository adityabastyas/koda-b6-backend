package repository

import (
	"context"
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `SELECT product_id, kategory_id, name, description, price, image_url FROM products`

	rows, err := lib.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ProductID, &p.KategoryID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `SELECT product_id, kategory_id, name, description,price, image_url FROM products WHERE product_id = $1`

	row := lib.DB.QueryRow(context.Background(), query, id)

	var p models.Product
	err := row.Scan(&p.ProductID, &p.KategoryID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
	if err != nil {
		return nil, err
	}

	return &p, nil

}
