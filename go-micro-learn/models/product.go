package models

import (
	"go-mircoservice-learn/go-micro-learn/entity"
	"strconv"
)

func NewProduct(id int, name string) *entity.Product {
	return &entity.Product{
		ID:    id,
		Name:  name,
		Stock: 100,
	}
}
func NewProductList(n int) *[]*entity.Product {
	res := make([]*entity.Product, n)
	for i := 0; i < n; i++ {
		res[i] = NewProduct(i+1, "name"+strconv.Itoa(i+1))
	}
	return &res
}
