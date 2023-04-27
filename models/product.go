package model

type Product struct {
	Name  string `json:"name"`
	Site  string `json:"site"`
	User  string `json:"user"`
	Price int64  `json:"price"`
}
