package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/AM07/backend/models"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Products)
}
