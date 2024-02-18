package repository

import (
	"ak/database"
	"ak/domain"
	"ak/models"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ADD PRODUCT ..

func AddProduct(product domain.Products) (domain.Products, error) {

	var p models.ProductReceiver

	if product.Quantity < 0 {
		return domain.Products{}, models.QuantityIsLessThanZero

	}

	if product.Price < 0 {
		return domain.Products{}, models.PriceIsLessThanZero
	}

	var count int

	errorRes := database.DB.Raw("select count (*) from products where name = ?", product.Name).Scan(&count)

	if errorRes.Error != nil {
		return domain.Products{}, errorRes.Error
	}

	if count > 0 {
		return domain.Products{}, models.PriceIsLessThanZero
	}

	err := database.DB.Raw("INSERT INTO products (name, sku, category_id, design_description, brand_id, quantity, price, product_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING name, sku, category_id, design_description, brand_id, quantity, price, product_status", product.Name, product.SKU, product.CategoryID, product.DesignDescription, product.BrandID, product.Quantity, product.Price, product.ProductStatus).Scan(&p).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {

			return domain.Products{}, errors.New("alerady exist ")

		}

		return domain.Products{}, err
	}

	var ProductResponse domain.Products
	err = database.DB.Raw("select * from products where products.name=?", p.Name).Scan(&ProductResponse).Error

	if err != nil {
		return domain.Products{}, err

	}

	return ProductResponse, nil
}

// ADD CATEGORY FROM PRODUCTS..

func AddCategory(category domain.Category) (domain.Category, error) {
	var b string
	check := database.DB.Raw("SELECT category_name FROM categories WHERE category_name = ?", category.CategoryName).Scan(&b).Error
	if check != nil {
		return domain.Category{}, check
	}
	if b != "" {
		return domain.Category{}, errors.New("name is already exist")

	}

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

	var existingCategory domain.Category
	if err := database.DB.Where("category_name = ?", new).First(&existingCategory).Error; err == nil {

		return domain.Category{}, errors.New("new category name already exists")
	}
	fmt.Println("exirt", existingCategory)
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

func CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := database.DB.Raw("select count(*) from categories where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			return errors.New("there is no category id")
		}

		if count < 1 {

			return errors.New("genre does not exist")
		}
	}
	return nil
}
func GetProductFromCategory(id int) ([]models.ProductBrief, error) {

	var product []models.ProductBrief
	err := database.DB.Raw(`
		SELECT *
		FROM products
		JOIN categories ON products.category_id = categories.id
		 where categories.id = ?
	`, id).Scan(&product).Error

	if err != nil {
		return []models.ProductBrief{}, err
	}

	if len(product) == 0 {
		return []models.ProductBrief{}, errors.New("there is no product have in these category")
	}

	fmt.Println("product", product)
	return product, nil
}

func GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := database.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {

	if page == 0 {

		page = 1
	}

	offset := (page - 1) * count

	var ProductBrief []models.ProductBrief

	err := database.DB.Raw(`SELECT
    products.id ,
    products.name ,
    products.sku,
    categories.category_name,
    products.brand_id,
    products.quantity,
    products.price,
    products.product_status
FROM
    products
JOIN
    categories ON products.category_id = categories.id
	limit ? offset ?`,
		count, offset).Scan(&ProductBrief).Error

	if err != nil {

		return []models.ProductBrief{}, err

	}

	fmt.Println("producgt brief form repo", ProductBrief)

	return ProductBrief, nil

}
func GetAllCategory(page int, count int) ([]domain.Category, error) {

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * count

	var categoryBrief []domain.Category

	err := database.DB.Raw("select id,category_name from categories limit ? offset ?", count, offset).Scan(&categoryBrief).Error

	if err != nil {
		return []domain.Category{}, err
	}

	return categoryBrief, nil

}

func AddProductOffer(productOffer models.ProductOfferReceiver) error {

	if productOffer.OfferLimit <= 0 {

		return errors.New("offer limit must be greater than zero")
	}

	if productOffer.DiscountPercentage < 0 {
		return errors.New("discount price is less than zero ")
	}

	//check if the offer with the offer name already exist 	in the database

	var count int

	err := database.DB.Raw("select count(*) from product_offers where offer_name =? and product_id =?", productOffer.OfferName, productOffer.ProductID).Scan(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {

		return errors.New("the offer alreay exist")

	}

	count = 0

	err = database.DB.Raw("select count(*) from product_offers where product_id =?", productOffer.ProductID).Scan(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		err = database.DB.Exec("delete from product_offers where product_id = ?", productOffer.ProductID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	err = database.DB.Exec("INSERT INTO product_offers (product_id, offer_name, discount_percentage, start_date, end_date, offer_limit,offer_used) VALUES (?, ?, ?, ?, ?, ?, ?)", productOffer.ProductID, productOffer.OfferName, productOffer.DiscountPercentage, startDate, endDate, productOffer.OfferLimit, 0).Error
	if err != nil {
		return err
	}

	return nil

}

func CheckProductPrice(productID int) (float64, error) {

	var price float64

	err := database.DB.Raw("SELECT price from products WHERE id = ?", productID).Scan(&price).Error
	if err != nil {
		return 0.0, err
	}

	return price, nil
}
