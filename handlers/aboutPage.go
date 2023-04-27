package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/AM07/backend/models"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About page")
	b := model.Server{URL: "http://localhost", Port: 8081}
	fmt.Fprintf(w, "Details of the server:\t")
	json.NewEncoder(w).Encode(b)
}
