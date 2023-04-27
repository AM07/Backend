package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AM07/backend/helpers"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(helpers.Products)
}
