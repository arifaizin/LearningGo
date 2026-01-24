package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99, Stock: 10},
	{ID: 2, Name: "Smartphone", Price: 499.99, Stock: 25},
	{ID: 3, Name: "Tablet", Price: 299.99, Stock: 15},
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Electronics", Description: "Electronic devices and gadgets"},
	{ID: 2, Name: "Home Appliances", Description: "Appliances for home use"},
	{ID: 3, Name: "Books", Description: "Various kinds of books"},
}

func main() {

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		case http.MethodPost:
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newProduct.ID = len(products) + 1
			products = append(products, newProduct)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newProduct)
		}
	})
	//get product by id /api/products/{id}
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/products/")

		if id == "" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			for _, product := range products {
				if fmt.Sprint(product.ID) == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(product)
					return
				}
			}
			http.NotFound(w, r)
		case http.MethodPut:
			var updatedProduct Product
			err := json.NewDecoder(r.Body).Decode(&updatedProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			for i, product := range products {
				if fmt.Sprint(product.ID) == id {
					products[i] = updatedProduct
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(updatedProduct)
					return
				}
			}
			http.NotFound(w, r)
		case http.MethodDelete:
			for i, product := range products {
				if fmt.Sprint(product.ID) == id {
					products = append(products[:i], products[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.NotFound(w, r)
		}
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK", "message": "Service is healthy"})
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		case http.MethodPost:
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newCategory)
		}
	})
	//get category by id /categories/{id}
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/categories/")

		if id == "" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			for _, category := range categories {
				if fmt.Sprint(category.ID) == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(category)
					return
				}
			}
			http.NotFound(w, r)
		case http.MethodPut:
			var updatedCategory Category
			err := json.NewDecoder(r.Body).Decode(&updatedCategory)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			for i, category := range categories {
				if fmt.Sprint(category.ID) == id {
					categories[i] = updatedCategory
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(updatedCategory)
					return
				}
			}
			http.NotFound(w, r)
		case http.MethodDelete:
			for i, category := range categories {
				if fmt.Sprint(category.ID) == id {
					categories = append(categories[:i], categories[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.NotFound(w, r)
		}
	})

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
