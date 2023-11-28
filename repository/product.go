package repository

import (
	"ak/database"
	"ak/domain"
	"ak/models"
	"errors"
	"fmt"
	"strconv"
)

// ADD PRODUCT ..

func AddProduct(product domain.Products) (domain.Products, error) {

	var p models.ProductReceiver

	// err := database.DB.Raw("insert into products (name,sku,category_id,design_descrition,brand_id,quantity,price,product_status,) values (?,?,?,?,?,?,?,?) returning name,sku,category_id,design_description,brand_id,quantity,price,product_status", product.Name, product.SKU, product.CategoryID, product.DesignDescription, product.BrandID, product.Quantity, product.Price, product.ProductStatus).Scan(&p).Error
	err := database.DB.Raw("INSERT INTO products (name, sku, category_id, design_description, brand_id, quantity, price, product_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING name, sku, category_id, design_description, brand_id, quantity, price, product_status", product.Name, product.SKU, product.CategoryID, product.DesignDescription, product.BrandID, product.Quantity, product.Price, product.ProductStatus).Scan(&p).Error

	if err != nil {
		return domain.Products{}, err
	}

	var ProductResponse domain.Products
	err = database.DB.Raw("select * from products where products.name=?", p.Name).Scan(&ProductResponse).Error

	// fmt.Println("Prodcuc name is ",p.Name)

	if err != nil {
		return domain.Products{}, err

	}

	return ProductResponse, nil
}

// ADD CATEGORY FROM PRODUCTS..

func AddCategory(category domain.Category) (domain.Category, error) {
	var b string
	err := database.DB.Raw("insert into categories (category_name) values (?) returning category_name", category.CategoryName).Scan(&b).Error
	if err != nil {
		fmt.Println("'vvv")
		return domain.Category{}, err
	}
	var categoryResponse domain.Category
	err = database.DB.Raw("SELECT id ,category_name FROM categories WHERE category_name = ?", b).Scan(&categoryResponse).Error
	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil
}

// CHECKING THE CATEGORY FROM REPO

func CheckCategoryExist(current string) (bool, error) {

	var i int

	if err := database.DB.Raw("select count(*) from categories where category_name =? ", current).Scan(&i).Error; err != nil {
		return false, errors.New("category name  is inavlid ")

	}

	if i < 1 {
		return false, errors.New("category name is not exist on database ")

	}
	return true, nil

}

//UPDATE CATEGORY FROM REPOSITORY

func UpdateCategory(current string, new string) (domain.Category, error) {

	if database.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")

	}

	if err := database.DB.Exec("update categories set category_name =$1 where category_name=$2", new, current).Error; err != nil {
		return domain.Category{}, err
	}

	var newupdate domain.Category

	if err := database.DB.First(&newupdate, "category_name=?", new).Error; err != nil {
		return domain.Category{}, err
	}

	return newupdate, nil

}

func DeleteCategory(categoryId string) error {

	id, err := strconv.Atoi(categoryId)
	fmt.Println("category id :", categoryId)

	if err != nil {
		return errors.New("couldn't convert")

	}

	result := database.DB.Exec("DELETE FROM categories where id =?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records are there")
	}
	return nil
}

// CHECKING  THERE IS A PRODUCT ID FOR UPDATING

func CheckProductExist(pid int) (bool, error) {
	var k int

	err := database.DB.Raw("select count(*)from products where id=?", pid).Scan(&k).Error
	if err != nil {

		return false, err
	}
	fmt.Println("repository proudct is ", k)

	if k == 0 {
		return false, errors.New("repositary dont have a product id")
	}
	return true, err
}

// UPDATE PRODUCT ...

func UpdateProduct(pid int, quantity int) (models.ProductUpdateReciever, error) {

	//
	if database.DB == nil {
		return models.ProductUpdateReciever{}, errors.New("database connection is nil")

	}

	if err := database.DB.Exec("UPDATE products SET quantity = quantity + $1 WHERE id= $2", quantity, pid).Error; err != nil {
		fmt.Println("quantity from update product ", quantity)
		return models.ProductUpdateReciever{}, err

	}

	var newdetails models.ProductUpdateReciever
	var newquantity int

	if err := database.DB.Raw("select quantity from products where id =?", pid).Scan(&newquantity).Error; err != nil {
		fmt.Println("pid was", pid)
		return models.ProductUpdateReciever{}, err
	}

	var pro = pid

	if pro == 0 {
		fmt.Println("there is nothing in pid")
		return models.ProductUpdateReciever{}, errors.New("pid nothing")

	}

	fmt.Println("pid was", pid)
	newdetails.ProductID = pid
	newdetails.Quantity = newquantity

	return newdetails, nil

}

// DELETE PRODUCT..

func DeleteProduct(productId string) error {

	id, err := strconv.Atoi(productId)

	if err != nil {
		return errors.New("coud'nt convert ")

	}

	result := database.DB.Exec("delete from products where id =$1", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that exist")
	}
	return nil
}
