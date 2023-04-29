package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AM07/backend/handlers"
	"github.com/AM07/backend/helpers"
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Site available at http://localhost:8081...")
}

func router() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Homepage).Methods("GET").Host("localhost").Schemes("http")
	r.HandleFunc("/about", handlers.AboutPage).Methods("GET")
	r.HandleFunc("/GetAllProducts", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/GetOneProduct", handlers.GetOneProduct).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}

func main() {
	helpers.AddEntry()
	F, err := os.OpenFile("error.logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer F.Close()

	log.SetOutput(F)
	router()
}
