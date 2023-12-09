package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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

	if err != nil {

		erres := response.ClientResponse(http.StatusBadGateway, "server error from usecase", nil, err.Error())
		c.JSON(http.StatusBadGateway, erres)
		return
	}

	successres := response.ClientResponse(http.StatusCreated, "succesed login user", LogedUser, nil)

	c.JSON(http.StatusBadGateway, successres)

}

func GetAllAddress(c *gin.Context) {
	user_id, _ := c.Get("user_id")
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

func AddAddress(c *gin.Context) {
	user_id, _ := c.Get("user_id")

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

		errREs := response.ClientResponse(http.StatusBadGateway, "error from adding address", nil, err)
		c.JSON(http.StatusBadGateway, errREs)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "added address sucessfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

//USER DETAILS FROM HANDLER

func UserDetails(c *gin.Context) {

	user_id, _ := c.Get("user_id")

	userdetails, err := usecase.UserDetails(user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "failed to retrive data", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "usre details", userdetails, nil)
	c.JSON(http.StatusOK, successRes)
}


func CheckOut(c *gin.Context) {

	userID,_ := c.Get("user_id")

	checkoutDetails,err := usecase.CheckOut(userID.(int))



	if err != nil {
		errRes:= response.ClientResponse(http.StatusInternalServerError,"failed to retrive",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
	}


	succesRes := response.ClientResponse(http.StatusOK,"chechout page loaded succesfully",checkoutDetails,nil)
	c.JSON(http.StatusOK,succesRes)	




}
