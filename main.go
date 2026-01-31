package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

var products = []models.Product{
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

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Success connected to database on:" + config.DatabaseURL)
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)

	//init db
	db, err := database.InitDB(config.DatabaseURL)
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}

	// fmt.Println("Starting server on :" + config.Port)
	// erro := http.ListenAndServe(":"+config.Port, nil)
	// if erro != nil {
	// 	fmt.Println("gagal running server", erro)
	// }

	// http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(products)
	// 	case http.MethodPost:
	// 		var newProduct models.Product
	// 		err := json.NewDecoder(r.Body).Decode(&newProduct)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 			return
	// 		}
	// 		newProduct.ID = len(products) + 1
	// 		products = append(products, newProduct)
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(newProduct)
	// 	}
	// })
	// //get product by id /api/products/{id}
	// http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
	// 	id := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// 	if id == "" {
	// 		http.NotFound(w, r)
	// 		return
	// 	}

	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		for _, product := range products {
	// 			if fmt.Sprint(product.ID) == id {
	// 				w.Header().Set("Content-Type", "application/json")
	// 				json.NewEncoder(w).Encode(product)
	// 				return
	// 			}
	// 		}
	// 		http.NotFound(w, r)
	// 	case http.MethodPut:
	// 		var updatedProduct models.Product
	// 		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 			return
	// 		}
	// 		for i, product := range products {
	// 			if fmt.Sprint(product.ID) == id {
	// 				products[i] = updatedProduct
	// 				w.Header().Set("Content-Type", "application/json")
	// 				json.NewEncoder(w).Encode(updatedProduct)
	// 				return
	// 			}
	// 		}
	// 		http.NotFound(w, r)
	// 	case http.MethodDelete:
	// 		for i, product := range products {
	// 			if fmt.Sprint(product.ID) == id {
	// 				products = append(products[:i], products[i+1:]...)
	// 				w.WriteHeader(http.StatusNoContent)
	// 				return
	// 			}
	// 		}
	// 		http.NotFound(w, r)
	// 	}
	// })
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
}
