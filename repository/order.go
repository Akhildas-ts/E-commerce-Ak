package repository

import (
	"ak/database"
	"ak/domain"
	"ak/models"
	"fmt"
)

func DoesCartExist(userID int) (bool, error) {

	var Exist bool

	err := database.DB.Raw("select exists(select 1 from carts where user_id= ?)", userID).Scan(&Exist).Error

	if err != nil {

		return false, err
	}

	fmt.Println("cart exist:",Exist)

	return Exist, nil
}

func AddreesExist(orderbody models.OrderIncoming) (bool, error) {

	var count int

	if err := database.DB.Raw("select count(*)from addresses where user_id =? and id =?", orderbody.UserID, orderbody.AddressID).Scan(&count).Error; err != nil {

		return false, err
	}

	return count > 0, nil
}

// func GetAllItemsFromCart(userid int) ([]models.Cart,error) {

// 	var count int

// 	var countResponse []models.Cart

// 	err := database.DB.Raw("select count(*)from carts where user_id=?",userid).Scan(&count).Error

// }

func CreateOrder(orderDetails domain.Order) error {

	err := database.DB.Create(&orderDetails).Error
	if err != nil {

		fmt.Println("error from the creae order from repo :")
		return err

		
	}
	return nil

}

func AddOrderItems(orderItemDetails domain.OrderItem, UserID int, ProductID uint, Quantity float64) error {

	// after creating the order delete all cart items and also update the quantity of the product
	err := database.DB.Omit("id").Create(&orderItemDetails).Error
	if err != nil {

		fmt.Println("inserted items::",err)

		fmt.Println("error from orderitems tables.: ")
		return err
	}

	err = database.DB.Exec("delete from carts where user_id = ? and product_id = ?", UserID, ProductID).Error
	if err != nil {
		return err
	}

	err = database.DB.Exec("update products set quantity = quantity - ? where id = ?", Quantity, ProductID).Error
	if err != nil {
		return err
	}

	return nil

}

func GetBriefOrderDetails(orderID string) (domain.OrderSuccessResponse, error) {

	var orderSuccessResponse domain.OrderSuccessResponse
	database.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderID).Scan(&orderSuccessResponse)
	return orderSuccessResponse, nil

}

//<<<<<<< CANCCEL THE ORDER <<<

func UserOrderRelationShip(orderid string ,userid int ) (int,error) {

	var testUserid int 

	err := database.DB.Raw("select user_id from orders where order_id=?",orderid).Scan(&testUserid).Error

	if err != nil{

		return -1,err
	}

	return testUserid,nil
}

func GetProductDetailsFromOrder(orderid string)([]models.OrderProducts,error) {

	var orderProductDetails []models.OrderProducts

	err := database.DB.Raw("select product_id,quantity from order_items where order_id=?",orderid).Scan(&orderProductDetails).Error

	if err != nil {
		return []models.OrderProducts{},err
	}

	return orderProductDetails,nil
}

func GetShipmentStatus(orderid string)(string,error){

	var shipmentstatus string

	err := database.DB.Raw("select shipment_status from orders where order_id=?",orderid).Scan(&shipmentstatus).Error

	if err != nil {
		return "",err
	}

	return shipmentstatus,nil
}


func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var orderDetails []models.OrderDetails
	fmt.Println("userid is", userId, "page is ", page, "count is ", count, "offset is", offset)
	database.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where user_id = ? limit ? offset ? ", userId, count, offset).Scan(&orderDetails)
	fmt.Println("order details is ", orderDetails)

	var fullOrderDetails []models.FullOrderDetails
	// for each order select all the associated products and their details
	for _, od := range orderDetails {

		var orderProductDetails []models.OrderProductDetails
		database.DB.Raw("select order_items.product_id,products.name as product_name,order_items.quantity,order_items.total_price from order_items inner join products on order_items.product_id = products.id where order_items.order_id = ?", od.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: od, OrderProductDetails: orderProductDetails})

	}

	return fullOrderDetails, nil

}

func CancelOrders(orderid string) error {

  shipmentstatus:= "cancelled"

  err := database.DB.Exec("update orders set shipment_status = ? where order_id =?",shipmentstatus,orderid).Error

  if err != nil {

	return err
  }
//   var paymentMethod int
	// err = database.DB.Raw("select payment_method_id from orders where order_id = ?", orderid).Scan(&paymentMethod).Error
	// if err != nil {
	// 	return err
	// }
	// if paymentMethod == 3 || paymentMethod == 2 {
	// 	err = database.DB.Exec("update orders set payment_status = 'refunded'  where order_id = ?", orderid).Error
	// 	if err != nil {
	// 		return err
	// 	}
	
	return nil

 
}

func UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {

	for _, od := range orderProducts {

		var quantity int
		if err := database.DB.Raw("select quantity from products where id = ?", od.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}

		od.Quantity += quantity
		if err := database.DB.Exec("update products set quantity = ? where id = ?", od.Quantity, od.ProductId).Error; err != nil {
			return err
		}
	}
	return nil

}


func CartExist(userID int) (bool, error) {
	var count int
	if err := database.DB.Raw("select count(*) from carts where user_id = ? ", userID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil

}

func DisplayCart(userID int) ([]models.Cart, error) {

	var count int
	if err := database.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := database.DB.Raw("select carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}

	return cartResponse, nil

}

// func GetAllAddress(userId int) (models.AddressInfoResponse, error) {
// 	var addressInfoResponse models.AddressInfoResponse
// 	if err := database.DB.Raw("select * from addresses where user_id = ?", userId).Scan(&addressInfoResponse).Error; err != nil {
// 		return models.AddressInfoResponse{}, err
// 	}
// 	return addressInfoResponse, nil
// }