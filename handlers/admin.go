package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var adminmodel models.AdminLogin

	if err := c.ShouldBindJSON(&adminmodel); err != nil {
		erres := response.ClientResponse(http.StatusBadGateway, "formate is not correct (admin)", nil, err)
		c.JSON(http.StatusBadGateway, erres)
		return

	}

	admin,err := usecase.AdminLogin(adminmodel)

	if err != nil {

		erres:= response.ClientResponse(http.StatusBadGateway,"server error from admin use case",nil,err)
		c.JSON(http.StatusBadGateway,erres)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK,"admin login succes ",admin,nil)
	c.JSON(http.StatusOK,succesRes)


} 


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
