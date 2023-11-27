package usecase

import (
	"ak/domain"
	"ak/models"
	"ak/repository"
	"fmt"
)

func AddProduct(product domain.Products) (domain.Products, error) {

	ProductResponse, err := repository.AddProduct(product)
	if err != nil {

		return domain.Products{}, err

	}
	return ProductResponse, nil

}

func UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error) {

	result, err := repository.CheckProductExist(pid)
	if err != nil {
		fmt.Println("error from repo")
		return models.ProductUpdateReciever{}, err
	}

	if !result {
		return models.ProductUpdateReciever{}, err

	}

	editProduct, err := repository.UpdateProduct(pid, stock)

	if err != nil {
		fmt.Println("error from update product ")
		return models.ProductUpdateReciever{}, err
	}
	return editProduct, err
}

func AddCategory(category domain.Category) (domain.Category, error) {

	ProductResponse, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}

	return ProductResponse, nil
}

func DeleteProduct(productId string) error {
	err := repository.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}
