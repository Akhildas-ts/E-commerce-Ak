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
