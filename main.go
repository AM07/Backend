package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AM07/backend/handlers"
	"github.com/gorilla/mux"
)

func init() {
	router()
}

func router() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomePage).Methods("GET")
	r.HandleFunc("/about", handlers.AboutPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Site running at http://localhost:8081...")
}
