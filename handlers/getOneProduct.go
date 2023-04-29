package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/AM07/backend/models"
)

func GetOneProduct(w http.ResponseWriter, r *http.Request) {

	// params := mux.Vars(r)
	// for _, product := range helpers.Products {
	// 	if product.Name == params["category"] {
	// 		json.NewEncoder(w).Encode(product)
	// 		return
	// 	}
	// }

	params1 := r.URL.Query().Get("category")
	params2 := r.URL.Query().Get("brand")
	if params1 != "" && params2 != "" {
		for _, product := range model.Products {
			if product.Category == params1 && product.Brand == params2 {
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		printOut := fmt.Sprintf("No matching product found for category %s and brand %s", params1, params2)
		json.NewEncoder(w).Encode(printOut)
		return
	} else if params1 != "" && params2 == "" {
		for _, product := range model.Products {
			if product.Category == params1 {
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		printOut := fmt.Sprintf("No matching product found for category %s", params1)
		json.NewEncoder(w).Encode(printOut)
		return
	} else if params1 == "" && params2 != "" {
		for _, product := range model.Products {
			if product.Brand == params2 {
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		printOut := fmt.Sprintf("No matching product found for %s", params2)
		json.NewEncoder(w).Encode(printOut)
		return
	} else {
		log.Panic("Missing category query parameter")
		return
	}
}
