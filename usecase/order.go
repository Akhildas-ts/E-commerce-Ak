package usecase

import (
	"ak/domain"
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
	"time"

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
		return domain.OrderSuccessResponse{}, models.CartEmpty
	}

	addressExist, err := repository.AddreesExist(orderbody)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, models.AddresNotFound
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
	discount_price, err := repository.GetCouponDiscountPrice(int(orderbody.UserID), orderDetail.GrandTotal)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	err = repository.UpdateCouponDetails(discount_price, orderDetail.UserID)

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderDetail.FinalPrice = orderDetail.GrandTotal - discount_price

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
		return models.UserNotMatch
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

		return errors.New("the order is  already  cancelled ,so no point to cancelling ")
	}

	// if shipmentStatus == "processing" {

	// 	return errors.New("the order is  already  cancelled ,so no point to cancelling ")
	// }
	err = repository.CancelOrders(orderid)

	if err != nil {

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
		return models.Invoice{}, models.CartEmpty
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

func OrderDelivered(orderID string) error {

	shipmentStatus, err := repository.GetShipmentStatus(orderID)
	if err != nil {
		fmt.Println("repo orderdeliverd")
	}

	fmt.Println("shipement status ", shipmentStatus)

	if shipmentStatus == "delivered" {

		return models.AlreadyReturn
	}

	if shipmentStatus == "processing" {
		shipmentStatus = "delivered"
		return repository.UpdateShipmentStatus(shipmentStatus, orderID)
	}

	return errors.New("order not placed or order id does not exist")
}

func ReturnOrder(orderid string) error {

	shipmentStatus, err := repository.GetShipmentStatus(orderid)

	if err != nil {

		return err
	}

	timeDelivered, err := repository.GetTimeDeliverdTime(orderid)

	if err != nil {

		return err
	}

	currentTime := time.Now()
	returnPeriod := timeDelivered.Add(time.Hour * 24 * 7)

	if currentTime.After(returnPeriod) {

		return models.CannotReturn

	}

	if shipmentStatus != "delivered" {
		return  models.ShipmentStatusIsNotDeliverd
	}

	if shipmentStatus == "return" {

		return models.AlreadyReturn
	}

	if shipmentStatus == "cancelled" {
		return models.AlreadyCancelled
	}

	if !currentTime.Before((returnPeriod)) {

		return errors.New("time is over ")
	}

	if shipmentStatus == "delivered" && currentTime.Before(returnPeriod) {

		shipmentStatus = "return"

		return repository.ReturnOrder(shipmentStatus, orderid)
	}

	return errors.New("can't return order ")
}

func GetOrderDetailsFromAdmin(page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := repository.GetOrderDetailsFromAdmin(page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil

}
