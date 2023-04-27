package helpers

import model "github.com/AM07/backend/models"

func AddEntry() {
	a := model.Product{Name: "laptop", User: "AMS", Site: "www.amazon.com", Price: 98000}
	Products = append(Products, a)
}
