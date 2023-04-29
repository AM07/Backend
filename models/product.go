package model

type Product struct {
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Site     string `json:"site"`
	User     string `json:"user"`
	Price    int64  `json:"price"`
}

var Products []Product
