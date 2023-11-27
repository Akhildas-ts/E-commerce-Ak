package handlers

import (
	"ak/domain"
	"ak/response"
	"ak/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ADD CATEEGORY

func AddCategory(c *gin.Context) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field provied in wrong formate ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	if err := validator.New().Struct(category); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "constrian are not satisfied ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	// if category.CategoryName == nil {

	// }

	addcategory, err := usecase.AddCategory(category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "use case formate is not correct  ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	succRes := response.ClientResponse(http.StatusOK, "Category was added succesfully", addcategory, nil)
	c.JSON(http.StatusOK, succRes)

}

//UPDATE CATEGORY ..

// func UpdateCategory(c *gin.Context){

// 	var uc models.Category

// 	if err := c.ShouldBindJSON(&uc);err != nil {
// 		errRes:= response.ClientResponse(http.StatusBadGateway,"field formate was not correct",nil,err)
// 		c.JSON(http.StatusBadGateway,errRes)
// 		return
// 	}

// 	 update,err := usecase.UpdateCategory(uc.ID,uc.Category)

// }
