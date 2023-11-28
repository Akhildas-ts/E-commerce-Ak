package handlers

import (
	"ak/domain"
	"ak/models"
	"ak/response"
	"ak/usecase"
	"fmt"
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

func UpdateCategory(c *gin.Context) {

	var uc models.SetNewName

	if err := c.ShouldBindJSON(&uc); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field formate was not correct", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	update, err := usecase.UpdateCategory(uc.Current, uc.New)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "update category did not exist ", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "successfully  update category", update, nil)
	c.JSON(http.StatusOK, succRes)

}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Query("id")

	fmt.Println("handler category id :", categoryId)

	err := usecase.DeleteCategory(categoryId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field provide has wrong formate", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "Delet record from categorie", nil, nil)
	c.JSON(http.StatusOK, succRes)
}
