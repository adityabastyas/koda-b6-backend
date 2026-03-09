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

func (r *ProductRepository) Create(input models.ProductInput) error {
	query := `INSERT INTO products (kategory_id, name, description, price, image_url) VALUES ($1, $2, $3, $4, $5)`

	_, err := lib.DB.Exec(context.Background(), query, input.KategoryID, input.Name, input.Description, input.Price, input.ImageURL)
	return err
}

func (r *ProductRepository) Update(id int, input models.ProductInput) error {
	query := `UPDATE products SET kategory_id=$1, name=$2, description=$3, price=$4, image_url=$5 WHERE product_id=$6`

	_, err := lib.DB.Exec(context.Background(), query, input.KategoryID, input.Name, input.Description, input.Price, input.ImageURL, id)
	return err
}
