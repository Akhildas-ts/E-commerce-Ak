package usecase

import (
	"ak/domain"
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

func OrderItemsFromCart(oderfromcart models.OrderFromCart, userId int) (domain.OrderSuccessResponse, error) {

	var orderbody models.OrderIncoming
	err := copier.Copy(&orderbody, &oderfromcart)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderbody.UserID = uint(userId)

	cartExist, err := repository.DoesCartExist(userId)
	if err != nil {
		return domain.OrderSuccessResponse{}, err

	}

	if !cartExist {
		return domain.OrderSuccessResponse{}, errors.New("cart empty can't order")
	}

	addressExist, err := repository.AddreesExist(orderbody)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("addres does not exist")
	}

	// GETTING ALL ITEMS FROM THE CART <<<<

	cartitems, err := repository.GetAllItemsFromCart(int(orderbody.UserID))

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	var orderDetail domain.Order
	var orderItemDetails domain.OrderItem

	// add general order details - that is to be added to orders table

	orderDetail = helper.CopyOrderDetails(orderDetail, orderbody)

	//get grand total iterarting through each products in carts

	for _, c := range cartitems {
		orderDetail.GrandTotal += c.TotalPrice
	}

	// if orderbody.PaymentID == 2 {
	// 	orderDetail.PaymentMethod ="not paid"
	// 	orderDetail.ShipmentStatus ="pending"
	// }

	err = repository.CreateOrder(orderDetail)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	for _, c := range cartitems {
		orderItemDetails.OrderID = orderDetail.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		err := repository.AddOrderItems(orderItemDetails, orderDetail.UserID, c.ProductID, c.Quantity)

		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}
	}

	orderSuccessResponse, err := repository.GetBriefOrderDetails(orderDetail.OrderId)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	return orderSuccessResponse, nil

}


func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil

}


func CancelOrder(orderid string, userid int) error {

	// <<<CHCHING THE ORDER IS DONE BY THE USER>>>>

	userTest, err := repository.UserOrderRelationShip(orderid, userid)

	if err != nil {

		return err
	}

	if userTest != userid {
		return errors.New("the order is done by the user )")
	}

	//>>GETTING PRODUCT ID FROM  ORDER'S<<<<<<

	orderProductDetails, err := repository.GetProductDetailsFromOrder(orderid)

	if err != nil {

		return err
	}

	shipmentStatus, err := repository.GetShipmentStatus(orderid)

	if err != nil {
		return err
	}

	if shipmentStatus == "delivered" {

		return errors.New("items already deliverd ,cannot cancel the order")
	}

	if shipmentStatus == "pending" || shipmentStatus == "returned" || shipmentStatus == "return" {

		messaage := fmt.Sprint(shipmentStatus)

		return errors.New("the order is " + messaage + ", so not   point the  cancelling")
	}

	if shipmentStatus == "cancelled" {
		fmt.Println("he")

		return errors.New("the order is  already  cancelled ,so no point to cancelling ")
	}

	if shipmentStatus == "processing" {

		return errors.New("the order is  already  cancelled ,so no point to cancelling ")
	}
	err = repository.CancelOrders(orderid)
	fmt.Println("eroor is ", err)

	if err != nil {
		fmt.Println("hey from cancel ordrs")
		return err
	}

	err = repository.UpdateQuantityOfProduct(orderProductDetails)
	if err != nil {
		return err
	}
	return nil

}

func ExecutePurchaseCOD(userID int, addressID int) (models.Invoice, error) {
	ok, err := repository.CartExist(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	if !ok {
		return models.Invoice{}, errors.New("cart doesnt exist")
	}
	cartDetails, err := repository.DisplayCart(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	addresses, err := repository.GetAllAddress(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	Invoice := models.Invoice{
		Cart:        cartDetails,
		AddressInfo: addresses,
	}
	return Invoice, nil

}
