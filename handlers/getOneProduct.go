package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AM07/backend/helpers"
)

func GetOneProduct(w http.ResponseWriter, r *http.Request) {

	// params := mux.Vars(r)
	// for _, product := range helpers.Products {
	// 	if product.Name == params["category"] {
	// 		json.NewEncoder(w).Encode(product)
	// 		return
	// 	}
	// }

	params := r.URL.Query().Get("category")
	if params != "" {
		for _, product := range helpers.Products {
			if product.Name == params {
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		printOut := fmt.Sprintf("No product found for %s", params)
		json.NewEncoder(w).Encode(printOut)
		return
	} else {
		log.Panic("Missing category query parameter")
	}
}
