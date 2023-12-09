package repository

import (
	"ak/database"
	"ak/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func QuantityOfProductInCart(user_id int, product_id int) (int, error) {

	var proudctQty int

	err := database.DB.Raw("select quantity from carts where user_id =? and product_id =?", user_id, product_id).Scan(&proudctQty).Error

	if err != nil {

		return 0, nil
	}

	return proudctQty, nil
}

func GetPriceOfProductFromID(product_id int) (float64, error) {

	var productPrice float64

	if err := database.DB.Raw("select price from products where id =?", product_id).Scan(&productPrice).Error; err != nil {
		fmt.Println("dont have product price")
		return 0.0, err
	}
	return productPrice, nil
}

func AddItemsIntoCart(user_id int, productid int, quantity int, prouductprice float64) error {

	if err := database.DB.Exec("insert into carts(user_id,product_id,quantity,total_price) values(?,?,?,?)", user_id, productid, quantity, prouductprice).Error; err != nil {
		return err
	}

	return nil
}

// func DisplayCart(userID int) ([]models.Cart, error) {

// 	var count int
// 	if err := database.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
// 		return []models.Cart{}, err
// 	}

// 	if count == 0 {
// 		return []models.Cart{}, nil
// 	}

// 	var cartResponse []models.Cart

// 	if err := database.DB.Raw("select carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
// 		return []models.Cart{}, err
// 	}

// 	return cartResponse, nil

// }

func ProductExist(userid int, productID int) (bool, error) {

	var count int
	if err := database.DB.Raw("select count(*) from carts where user_id =? and product_id =?", userid, productID).Scan(&count).Error; err != nil {

		fmt.Println("error from productExist")
		return false, err

	}

	if count > 0 {
		fmt.Println("there is no product ::::::")
	}

	return count > 0, nil
}

func GetQuantityAndProductDetails(userid int, productid int, cartdetails struct {
	Quantity   int
	TotalPrice float64
}) (struct {
	Quantity   int
	TotalPrice float64
}, error) {

	if err := database.DB.Raw("select quantity,total_price from carts where user_id =? and product_id =?", userid, productid).Scan(&cartdetails).Error; err != nil {
		return struct {
			Quantity   int
			TotalPrice float64
		}{}, err
	}

	return cartdetails, nil
}
func RemoveProductFromCart(userID int, product_id int) error {

	if err := database.DB.Exec("delete from carts where user_id =? and product_id =?", uint(userID), uint(product_id)).Error; err != nil {
		return err
	}

	return nil
}

func CartAfterRemovalOfProduct(user_id int) ([]models.Cart, error) {
	var cart []models.Cart
	if err := database.DB.Raw("select carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join products on carts.product_id = products.id where carts.user_id = ?", user_id).Scan(&cart).Error; err != nil {
		return []models.Cart{}, err
	}
	return cart, nil
}

func GetTotalPrice(userID int) (models.CartTotal, error) {

	var cartTotal models.CartTotal
	err := database.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	err = database.DB.Raw("select firstname as user_name from users where id = ?", userID).Scan(&cartTotal.UserName).Error
	if err != nil {
		return models.CartTotal{}, err
	}

	return cartTotal, nil

}

func DisplayCart(userid int) ([]models.Cart, error) {

	var count int

	if err := database.DB.Raw("select count(*) from  carts where user_id =?", userid).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := database.DB.Raw("select carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id =users.id inner join products on carts.product_id =products.id where user_id =? ", userid).First(&cartResponse).Error; err != nil {

		return []models.Cart{}, err
	}

	return cartResponse, nil
}

//GETTING ALL ITEMS FROM CARTS FOR THE CHECKOUT PAGE <<<<<<

func GetAllItemsFromCart(userid int) ([]models.Cart, error) {
	var count int

	var cartResponse []models.Cart

	err := database.DB.Raw("select count(*) from carts where user_id=?", userid).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		fmt.Println("there is no product from carts : count == 0")
		return []models.Cart{}, nil

	}

	err = database.DB.Raw("select carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userid).First(&cartResponse).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err

		}

		return []models.Cart{}, err
	}

	return cartResponse, nil
}
