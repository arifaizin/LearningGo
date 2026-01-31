package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ProductRepository struct {
	// You can add fields like DB connection here if needed
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	// Implement logic to get all products from the database
	var products []models.Product
	// Example query (adjust according to your database schema)
	rows, err := r.db.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
