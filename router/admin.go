package routes

import (
	"ak/handlers"
	"ak/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/admin", handlers.AdminLogin)

	//MIDDLE WARE 

	r.Use(middleware.AuthorizationMiddleware())
	{
		r.GET("/dashboard", handlers.DashBord)
	}

	//PRODUCT 
	product := r.Group("/products")
	{
		product.POST("", handlers.AddProduct)
		product.PUT("",handlers.UpdateProduct)
		product.DELETE("",handlers.DeleteProduct)
	}

	// Categorry

	category := r.Group("/category") 
	{
		category.POST("",handlers.AddCategory)
	}

	


	return r

}
