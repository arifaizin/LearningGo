package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ProductHandler struct {
	// You can add fields like DB connection here if needed
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProcuts GET /api/products & POST /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Implement logic to get all products
		h.GetAll(w, r)
	case http.MethodPost:
		// Implement logic to create a new product
		// h.Create(w, r)
	}
}

// HandleProduct GET /api/products/{id}
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
