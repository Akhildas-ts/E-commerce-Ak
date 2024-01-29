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
	if stock < 0 {
		return models.ProductUpdateReciever{},errors.New("quantiy is less than zero")
	}

	result, err := repository.CheckProductExist(pid)
	if err != nil {
		
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

	if new == "" {
		return domain.Category{},errors.New("new name cannot be nil ")
	}

	result, err := repository.CheckCategoryExist(current)
	if err != nil {

		return domain.Category{}, err
	}

	if !result {
		return domain.Category{}, err

	}

	newupdate, err := repository.UpdateCategory(current, new)

	if err != nil {

		return domain.Category{}, err

	}

	return newupdate, nil

}

func DeleteCategory(categoryId string) error {
	err := repository.DeleteCategory(categoryId)
	if err != nil {
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

func FilterCategory(data map[string]int) ([]models.ProductBrief, error) {

	err := repository.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	var productFromCategory []models.ProductBrief
	for _, id := range data {

		product, err := repository.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductBrief{}, err
		}
		for _, product := range product {   

		

			quantity, err := repository.GetQuantityFromProductID(product.ID)
			if err != nil {
				return []models.ProductBrief{}, err
			}
			if quantity == 0 {
				product.ProductStatus = "out of stock"
			} else {
				product.ProductStatus = "in stock"
			}
			if product.ID != 0 {
				productFromCategory = append(productFromCategory, product)
			}
		}

		

	}
	return productFromCategory, nil

}

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {

	productsBrief, err := repository.ShowAllProducts(page, count)

	if err != nil {

		fmt.Println("error from usecase")
		return []models.ProductBrief{}, err
	}

	for i := range productsBrief {

		p := &productsBrief[i]

		if p.Quantity == 0 {

			p.ProductStatus = "out of stock"
		} else {

			p.ProductStatus = "in stock"
		}

	}

	return productsBrief, nil

}

func GetAllCategory(page int, count int) ([]domain.Category, error) {

	categoryBreif, err := repository.GetAllCategory(page, count)
	if err != nil {
		return []domain.Category{}, err
	}

	return categoryBreif, nil

}
func AddProductOffer(productOffer models.ProductOfferReceiver) error {

	return repository.AddProductOffer(productOffer)
}
