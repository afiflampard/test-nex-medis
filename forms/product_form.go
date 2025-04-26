package forms

import "fmt"

type ProductForm struct{}

type ProductFormInput struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Price       float64 `form:"column:price" json:"price"`
	Stock       int     `form:"stock" json:"stock"`
	Status      string  `form:"status" json:"status"`
}

type ProductStatus struct {
	Status []string `form:"status" json:"status"`
}

func (p ProductForm) ValidatePrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("price cannot be lower than zero")
	}
	return nil
}
