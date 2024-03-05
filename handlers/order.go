package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Order Items From Cart
// @Description Order Items from cart
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param OrderFromCart body models.OrderFromCart true "Items Ordering From The Cart"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /order [post]
func OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get(models.User_id)

	userID := id.(int)

	var OrderFromCart models.OrderFromCart

	if err := c.ShouldBind(&OrderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "formate invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	ordersuccessResponse, err := usecase.OrderItemsFromCart(OrderFromCart, userID)

	if errors.Is(err, models.AddresNotFound) || errors.Is(err, models.CartEmpty) {
		errorRes := response.ClientResponse(http.StatusNotFound, "record not found", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return

	}

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", ordersuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Order Details to user side
// @Description Get all order details done by user to user side
// @Tags User Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "page number"
// @Param count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /order/{id} [get]
func GetOrderDetails(c *gin.Context) {

	pageStr := c.Param(models.Page)
	pages, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format,check it ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query(models.Count))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, _ := c.Get(models.User_id)
	userID := id.(int)

	fullOrderDetails, err := usecase.GetOrderDetails(userID, pages, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusNotFound, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// cannceling the order <<<<
// @Summary Cancel order
// @Description Cancel order by the user using order ID
// @Tags User Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /order/{id} [put]
func CancelOrder(c *gin.Context) {

	orderID := c.Param(models.ID)

	id, _ := c.Get(models.User_id)
	userID := id.(int)

	err := usecase.CancelOrder(orderID, userID)

	if errors.Is(err, models.UserNotMatch) {

		errorRes := response.ClientResponse(http.StatusBadRequest, "request not correct ", nil, err.Error())

		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadGateway, "failed to cannel the order ", nil, err.Error())

		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	succesesRes := response.SuccessClientResponse(http.StatusOK, "order is cancelled:::")
	c.JSON(http.StatusOK, succesesRes)
}


func PlaceOrder(c *gin.Context) {
	userID, _ := c.Get(models.User_id)
	userId := userID.(int)
	straddress := c.Param(models.AddressId)
	paymentMethod := c.Param(models.Payment)
	addressId, err := strconv.Atoi(straddress)

	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	if paymentMethod == "1" {

		Invoice, err := usecase.ExecutePurchaseCOD(userId, addressId)

		if errors.Is(err, models.CartEmpty) {

			errorRes := response.ClientResponse(http.StatusBadRequest, "record not found", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}

		if err != nil {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "error in making cod ", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}
		successRes := response.ClientResponse(http.StatusOK, "Placed Order with cash on delivery", Invoice, nil)
		c.JSON(http.StatusOK, successRes)
	}
}

// @Summary ORDER DELIVERD
// @Description Order deliverd from user side
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /delivered/{id} [get]
func OrderDelivered(c *gin.Context) {

	orderID := c.Param(models.Order_ID)
	err := usecase.OrderDelivered(orderID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "order could not be delivered", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusOK, "successfully delivered the product")
	c.JSON(http.StatusOK, successRes)

}

// @Summary RETURN ORDER
// @Description return order from the user side
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cancel/{id} [get]
func ReturnOrder(c *gin.Context) {

	orderId := c.Param(models.Order_ID)

	err := usecase.ReturnOrder(orderId)

	if errors.Is(err, models.CannotReturn) || errors.Is(err, models.AlreadyReturn) || errors.Is(err, models.AlreadyCancelled) || errors.Is(err, models.ShipmentStatusIsNotDeliverd) {

		errRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())

		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	if err != nil {

		errRes := response.ClientResponse(http.StatusInternalServerError, "order could not be returned", nil, err.Error())

		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusOK, "successfully returned")
	c.JSON(http.StatusOK, successRes)

}

// @Summary GET ORDER DETAILS FROM ADMIN
// @Description Order details from admin side
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path int true "page number"
// @Param count query int true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/order/{page} [get]
func GetOrderDetailsFromAdmin(c *gin.Context) {

	pageStr := c.Param(models.Page)
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query(models.Count))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	fullOrderDetails, err := usecase.GetOrderDetailsFromAdmin(page, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}
