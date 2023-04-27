package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/AM07/backend/models"
	"github.com/gorilla/mux"
)

var products []model.Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// http.ServeFile(w, r, "./static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About page")
	a := model.Product{Name: "laptop", User: "AMS", Site: "www.amazon.com", Price: 98000}
	products = append(products, a)
	fmt.Fprintf(w, "Details of the server:\t")
	json.NewEncoder(w).Encode(a)
}

func handlefunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/about", aboutPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlefunc()
}
