package usecase

import (
	"ak/domain"
	"ak/models"
	"ak/repository"
	"errors"
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

//UPDATE CATEGOREY ....

func UpdateCategory(current string, new string) (domain.Category, error) {

	result, err := repository.CheckCategoryExist(current)
	if err != nil {
		fmt.Println("from name ")
		return domain.Category{}, err
	}

	if !result {
		return domain.Category{}, errors.New("error from checkCategory exist ")

	}

	newupdate, err := repository.UpdateCategory(current, new)

	if err != nil {
		return domain.Category{}, errors.New("errors from update category ")

	}

	return newupdate, nil

}

func DeleteCategory(categoryId string)error{
	err := repository.DeleteCategory(categoryId)
	if err != nil{
		return err
	}
	return nil

}

func DeleteProduct(productId string) error {
	err := repository.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}
