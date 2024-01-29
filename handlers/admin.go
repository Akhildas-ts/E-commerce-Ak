package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
// @Summary Admin Login
// @Description Login handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param  admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/adminlogin [post]
func AdminLogin(c *gin.Context) {
	var adminmodel models.AdminLogin

	if err := c.ShouldBindJSON(&adminmodel); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "formate is not correct (admin)", nil, err)
		c.JSON(http.StatusBadGateway, erres)
		return

	}

	if err := validator.New().Struct(adminmodel); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "constrian are not satisfied ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	admin, err := usecase.AdminLogin(adminmodel)

	if err != nil {

		erres := response.ClientResponse(http.StatusBadGateway, "server error from admin use case", nil, err)
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "admin login succes ", admin, nil)
	c.JSON(http.StatusOK, succesRes)

}
// @Summary Admin Dashboard
// @Description Get Amin Home Page with Complete Details
// @Tags Admin Dash Board
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/dashboard [GET]
func DashBord(c *gin.Context) {

	adminDashBoard, err := usecase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed fine", adminDashBoard, nil)
	c.JSON(http.StatusOK, successRes)
}
// @Summary Approve Order
// @Description Approve Order from admin side which is in processing state
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order Id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/approve-order/{id} [get]
func ApproveOrder(c *gin.Context) {

	orderId := c.Param(models.Order_ID)

	err := usecase.ApproveOrder(orderId)

	if errors.Is(err, models.OrderIsAlreadyPlaced) || errors.Is(err,models.OrderIsAlreadyCancelled){
		errorRes := response.ClientResponse(http.StatusBadRequest, "Bad request", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not approve the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusOK, "Order approved successfully")
	c.JSON(http.StatusOK, successRes)

}
// @Summary Cancel Order Admin
// @Description Cancel Order from admin side
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/cancel-order/{id} [get]
func CancelOrderFromAdminSide(c *gin.Context) {

	orderID := c.Param(models.Order_ID)

	err := usecase.CancelOrderFromAdminSide(orderID)
	if errors.Is(err, models.OrderIsAlreadyCancelled) {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Bad request", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return

	}

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusOK, "Cancel Successfull")
	c.JSON(http.StatusOK, successRes)

}
