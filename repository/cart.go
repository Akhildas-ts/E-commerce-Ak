package repository

import (
	"ak/database"
	"ak/helper"
	"ak/models"
	"errors"
	"fmt"
	"strconv"

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

	var productCount int
	if err := database.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", productid).Scan(&productCount).Error; err != nil {

		return err
	}

	if productCount == 0 {
		return errors.New("product does not exist")
	}

	if err := database.DB.Exec("insert into carts(user_id,product_id,quantity,total_price) values(?,?,?,?)", user_id, productid, quantity, prouductprice).Error; err != nil {

		return err
	}

	return nil
}

func IncreaseQuantiyInCart(userID int, productId int, quantity int, productPrice float64) error {

	var cartdetails models.Cart

	err := database.DB.
		Table("carts").
		Where("user_id = ? AND product_id = ?", userID, productId).
		Order("product_id").
		Limit(1).
		First(&cartdetails).
		Error

	if err != nil {
		fmt.Println("same same")
		return err
	}
	newquantity := cartdetails.Quantity + float64(quantity)
	newTotalprice := float64(newquantity) * productPrice

	err = database.DB.
		Table("carts").
		Where("user_id = ? AND product_id = ?", userID, productId).
		Updates(map[string]interface{}{
			"quantity":    newquantity,
			"total_price": newTotalprice,
		}).
		Error
	if err != nil {
		return err
	}

	return nil

}

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

func UpdateCartDetails(cartDetails struct {
	Quantity   int
	TotalPrice float64
}, userID int, productId int) error {

	if err := database.DB.Exec("update carts set quantity = ?,total_price = ? where user_id = ? and product_id = ?", cartDetails.Quantity, cartDetails.TotalPrice, userID, productId).Error; err != nil {
		return err
	}

	return nil
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

	var discount_price float64

	discount_price, err = helper.GetCouponDiscountPrice(userID, cartTotal.TotalPrice)

	fmt.Println("discount price", discount_price)
	if err != nil {
		return models.CartTotal{}, err
	}

	formattedTotalPrice := fmt.Sprintf("%.2f", cartTotal.TotalPrice)
	cartTotal.TotalPrice, _ = strconv.ParseFloat(formattedTotalPrice, 64)
	cartTotal.FinalPrice = cartTotal.TotalPrice - discount_price
	formattedFinalPrice := fmt.Sprintf("%.2f", cartTotal.FinalPrice)
    cartTotal.FinalPrice, _ = strconv.ParseFloat(formattedFinalPrice, 64)


	fmt.Println("carttotal .finalprice", cartTotal.FinalPrice)

	return cartTotal, nil

}

func DiscountReason(userID int, tableName string, discountLabel string, discountApplied *[]string) error {

	var count int
	err := database.DB.Raw("select count(*) from "+tableName+" where used = false and user_id = ?", userID).Scan(&count).Error

	if err != nil {
		return err
	}

	if count != 0 {
		*discountApplied = append(*discountApplied, discountLabel)
		count = 0
	}
	fmt.Println("discount", discountApplied)

	return nil
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
		return []models.Cart{}, errors.New("there is no product from cart")

	}

	err = database.DB.Raw("select carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userid).First(&cartResponse).Error

	if len(cartResponse) == 0 {
		fmt.Println("len")
		return []models.Cart{}, err
	}

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			if len(cartResponse) == 0 {
				fmt.Println("len")
				return []models.Cart{}, err
			}
			return []models.Cart{}, err

		}

		return []models.Cart{}, err
	}

	return cartResponse, nil
}
func GetTotalPriceFromCart(userID int) (float64, error) {

	var totalPrice float64
	err := database.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}

	return totalPrice, nil
}
