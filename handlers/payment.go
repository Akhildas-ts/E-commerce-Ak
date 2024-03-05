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

func MakePaymentRazorPay(c *gin.Context) {

	orderID := c.Query(models.Order_ID)
	userID := c.Query(models.User_id)
	user_Id, _ := strconv.Atoi(userID)

	orderDetail, razorID, err := usecase.MakePaymentRazorPay(orderID, user_Id)
	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"final_price": orderDetail.FinalPrice * 100,
		"razor_id":    razorID,
		"user_id":     userID,
		"order_id":    orderDetail.OrderId,

		"total": int(orderDetail.FinalPrice),
	}) 

}

func VerifyPayment(c *gin.Context) {
	orderID := c.Query(models.Order_ID)
	paymentID := c.Query(models.Payment_id)

	err := usecase.SavePaymentDetails(paymentID, orderID)

	if errors.Is(err,models.AlreadyPaid){
		errorRes := response.ClientResponse(http.StatusBadRequest, "request is not correct", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return

	}

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.SuccessClientResponse(http.StatusOK, "Successfully updated payment detailse")
	c.JSON(http.StatusOK, successRes)
}
