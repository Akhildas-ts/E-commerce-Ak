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

//ADD PRODUCT

func AddProduct(c *gin.Context) {

	var product domain.Products

	if err := c.ShouldBindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field provide are wrong formate (product )", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	// if product.Category.CategoryName == nil{

	// }

	err := validator.New().Struct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "Product structer is not correct ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	ProductResponse, err := usecase.AddProduct(product)
	fmt.Println("product", ProductResponse)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "can't add the product", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, " Product is addeed ", ProductResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

//UPDATE THE PRODUCT*****

func UpdateProduct(c *gin.Context) {

	var p models.ProductUpdate

	if err := c.ShouldBindJSON(&p); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field formate provide wrong formate", nil, err)

		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	updating, err := usecase.UpdateProduct(p.ProductId, p.Quantity)

	// if updating.ProductID == 0 {
	// 	fmt.Println("id onulley")
	// 	return
	// }

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not update the product ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updating the product ", updating, nil)
	c.JSON(http.StatusOK, successRes)

}

/// DELET THE PROUDCT ****

func DeleteProduct(c *gin.Context) {
	productID := c.Query("id")

	if productID == "" {

		fmt.Println("product id is nil ")
		return

	}

	err := usecase.DeleteProduct(productID)
	if err != nil {
		fmt.Println("product id :", productID)

		errRes := response.ClientResponse(http.StatusBadGateway, "field provide are wrong formate ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "Product was deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, succRes)

}
