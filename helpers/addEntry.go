package helpers

import model "github.com/AM07/backend/models"

func AddEntry() {
	a := model.Product{Category: "laptop", Brand: "HP", User: "AMS", Site: "www.amazon.com", Price: 98000}
	model.Products = append(model.Products, a)
}
