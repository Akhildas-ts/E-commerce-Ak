package handlers

import (
	"ak/domain"
	"ak/models"
	"ak/response"
	"ak/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ADD CATEGORY
// @Summary ADD CATEGORY
// @Description Add a new Category for products
// @Tags Admin category
// @Accept json
// @Produce json
// @Security Bearer
// @Param category body domain.Category true "Add new Category "
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category [post]
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

// UPDATE CATEGORY
// @Summary UPDATE CATEGORY
// @Description UPDATE category from product
// @Tags Admin category
// @Accept json
// @Produce json
// @Security Bearer
// @Param category body models.SetNewName true "Update category "
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category [put]
func UpdateCategory(c *gin.Context) {

	var uc models.SetNewName

	if err := c.ShouldBindJSON(&uc); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field formate was not correct", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	update, err := usecase.UpdateCategory(uc.Current, uc.New)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "update category did not exist ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "successfully  update category", update, nil)
	c.JSON(http.StatusOK, succRes)

}

// Delete Category
// @Summary DELETE CATEGORY
// @Description Add a new Category for products
// @Tags Admin category
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query string true "Category ID to be deleted"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category [delete]
func DeleteCategory(c *gin.Context) {
	categoryId := c.Query(models.ID)

	err := usecase.DeleteCategory(categoryId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provide has wrong formate", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "Delet record from categorie", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// Get Category
// @Summary GET CATEGORY
// @Description Get all category
// @Tags Admin category
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path int true "Page number"
// @Param count query int false "Page count" default(10)
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category/{page} [get]
func GetAllCategory(c *gin.Context) {

	pageStr := c.Param(models.Page)

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query(models.Count))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	category, err := usecase.GetAllCategory(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "can't get the cateogry details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfullu get the categories", category, nil)

	c.JSON(http.StatusOK, succesRes)

}
