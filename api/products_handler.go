package api

import "github.com/typetokrish/simple-go-api/model"

type ProductsApi interface {
	GetProducts() []model.Product
}

type ProductsHandler struct {
}

func NewProductsHandler() ProductsApi {
	return &ProductsHandler{}
}

func (p *ProductsHandler) GetProducts() []model.Product {

	products := []model.Product{
		{
			ID:          1,
			Name:        "macos catalina",
			Description: "Mac OS for development",
			Price:       100.69,
		},
		{
			ID:          2,
			Name:        "samsung galaxy",
			Description: "Top mobile phone with quality camera",
			Price:       60.69,
		},
	}

	return products
}
