package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {

	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		errres := response.ClientResponse(http.StatusBadGateway, "(otp)- field provied wrong formate", nil,err.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}

	err := usecase.SendOTP(phone.PhoneNumber)

	if err != nil {

		errores := response.ClientResponse(http.StatusBadGateway, "can't send otp", nil, err.Error())
		c.JSON(http.StatusBadGateway, errores)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "Otp send succesfull", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}

func VerifyOTP(c *gin.Context) {
	var code models.VerifyData
	if err := c.ShouldBindJSON(&code); err != nil {
		errres := response.ClientResponse(http.StatusBadGateway, "json format is not correct", nil, err.Error())
		c.JSON(http.StatusBadGateway, errres)
		return
	}

	users, err := usecase.VerifyOTP(code)
	if err != nil {
		errres := response.ClientResponse(http.StatusBadGateway, "could Not verify the otp ", nil, err)
		c.JSON(http.StatusBadGateway, errres)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Otp Verification Done", users, nil)
	c.JSON(http.StatusOK, successRes)

}
