package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/AM07/backend/models"
)

var products []model.Product

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// http.ServeFile(w, r, "./static/home.html")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About page")
	a := model.Product{Name: "laptop", User: "AMS", Site: "www.amazon.com", Price: 98000}
	products = append(products, a)
	fmt.Fprintf(w, "Details of the server:\t")
	json.NewEncoder(w).Encode(a)
}
