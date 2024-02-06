package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary SignUp functionality for user
// @Description SignUp functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.SignupDetail true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /signup [post]
func Signup(c *gin.Context) {

	var usersign models.SignupDetail

	if err := c.ShouldBindJSON(&usersign); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong formattttt 🙌", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	// CHEKING THE DATA ARE SENDED IN CORRECT FORMET OR NOT

	if err := validator.New().Struct(usersign); err != nil {

		errres := response.ClientResponse(404, "They are not in format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errres)
		return
	}

	usercreate, err := usecase.UsersingUp(usersign)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "user signup format error ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User sign up succsed", usercreate, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary LogIn functionality for user
// @Description LogIn functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /login [post]
func UserLoginWithPassword(c *gin.Context) {

	var LoginUser models.UserLogin

	if err := c.ShouldBindJSON(&LoginUser); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "Login field provided in wrong way ", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return

	}

	////////

	if err := validator.New().Struct(LoginUser); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "Login field was wrong formate ahn", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	LogedUser, err := usecase.UserLogged(LoginUser)
	if errors.Is(err, models.ErrEmailNotFound) {

		erres := response.ClientResponse(http.StatusBadRequest, "invalid email", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	if err != nil {

		erres := response.ClientResponse(500, "server error from usecase", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	successres := response.ClientResponse(http.StatusCreated, "succesed login user", LogedUser, nil)

	c.JSON(http.StatusOK, successres)

}

// @Summary Gett All Address
// @Description From Gett all Address from User
// @Tags User User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /showaddres [get]
func GetAllAddress(c *gin.Context) {
	user_id, _ := c.Get(models.User_id)
	addressInfo, err := usecase.GetAllAddress(user_id.(int))

	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadGateway, "failed to retrive the details", nil, err.Error())
		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "User address", addressInfo, nil)
	c.JSON(http.StatusOK, succRes)
}

// ADD ADDRESS FROM HADNLER

// @Summary AddAddress functionality for user
// @Description AddAddress functionality at the user side
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /addaddress [post]
func AddAddress(c *gin.Context) {
	fmt.Println("hey fom")
	user_id, _ := c.Get(models.User_id)

	var address models.AddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(address)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints does not match", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := usecase.Addaddress(user_id.(int), address); err != nil {

		errREs := response.ClientResponse(http.StatusBadRequest, "error from adding address", nil, err.Error())
		c.JSON(http.StatusBadGateway, errREs)
		return
	}
	

	successRes := response.SuccessClientResponse(http.StatusOK, "added address sucessfully")
	c.JSON(http.StatusOK, successRes)
}

//USER DETAILS FROM HANDLER

// @Summary User Details
// @Description User Details from User Profile
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /userdetails [get]
func UserDetails(c *gin.Context) {

	user_id, _ := c.Get(models.User_id)

	userdetails, err := usecase.UserDetails(user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "failed to retrive data", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "usre details", userdetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary CheckOut Page
// @Description CheckOut page from user
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /checkout [get]
func CheckOut(c *gin.Context) {

	userID, _ := c.Get(models.User_id)

	checkoutDetails, err := usecase.CheckOut(userID.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to retrive", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "chechout page loaded succesfully", checkoutDetails, nil)
	c.JSON(http.StatusOK, succesRes)

}

// @Summary Apply referrals
// @Description Apply referrals amount to order
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /referral/apply [get]
func ApplyReferral(c *gin.Context) {

	userId, _ := c.Get(models.User_id)

	message, err := usecase.ApplyReferral(userId.(int))

	if errors.Is(err, models.CartEmpty) {

		errRes := response.ClientResponse(http.StatusNotFound, "record not found", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return

	}

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add referral amount", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	if message != "" {
		errRes := response.ClientResponse(http.StatusOK, message, nil, nil)
		c.JSON(http.StatusOK, errRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusOK, "successfully added referral amount")
	c.JSON(http.StatusOK, successRes)

}
