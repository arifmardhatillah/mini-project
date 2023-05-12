package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type ProductUsecase interface {
	CreateProduct(product *model.Order) error
	GetProduct(id uint) (product model.Product, err error)
	GetListProducts() (Products []model.Product, err error)
	UpdateProduct(product *model.Product) (err error)
	DeleteProduct(id uint) (err error)
}

func CreateProduct(name, description string, price, stock uint) (*model.Product, error) {
	products, err := database.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, v := range products {
		if v.Name == name {
			return nil, errors.New("product already exist")
		}
	}
	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	err = database.CreateProduct(product)
	if err != nil {
		fmt.Println("error creating product")
		return nil, err
	}
	return product, nil
}

func GetProduct(id uint) (product model.Product, err error) {
	product, err = database.GetProduct(id)
	if err != nil {
		fmt.Println("GetProduct: Error getting product from database")
		return
	}
	return
}

func GetListProducts() (products []model.Product, err error) {
	products, err = database.GetProducts()
	if err != nil {
		fmt.Println("GetListProducts : Error getting product from database")
		return
	}
	return
}

func UpdateProduct(product *model.Product, stock uint) (err error) {
	product.Stock = stock

	err = database.UpdateProduct(product)
	if err != nil {
		fmt.Println("UpdateProduct: Error updating product, err:", err)
		return
	}

	return
}

func DeleteProduct(id uint) (err error) {
	product := model.Product{}
	product.ID = id
	err = database.DeleteProduct(&product)
	if err != nil {
		fmt.Println("DeleteProduct : error deleting product, err: ", err)
		return
	}

	return
}
