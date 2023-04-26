package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Purpose string `json:"purpose"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// http.ServeFile(w, r, "./static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "About page")
	a := Server{Name: "MyWebsite", Owner: "AMS", Purpose: "Backend development"}
	fmt.Fprintf(w, "Details of the server:\t")
	json.NewEncoder(w).Encode(a)
}

func handlefunc() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	// log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlefunc()
	http.ListenAndServe(":8081", nil)
}
