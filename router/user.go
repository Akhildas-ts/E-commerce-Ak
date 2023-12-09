package routes

import (
	"ak/handlers"
	"ak/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)
	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	r.Group("/users")
	r.Use(middleware.AuthMiddleware())
	{
		r.GET("/showaddres", handlers.GetAllAddress)
		r.POST("/addaddress", handlers.AddAddress)
		r.GET("/userdetails", handlers.UserDetails)

		// CART
		cart := r.Group("/cart")

		{
			cart.POST("/addtocart/:product_id", handlers.AddToCart)
			cart.DELETE("/removefromcart/:product_id", handlers.RemoveFromCart)
			cart.GET("/displaycart", handlers.DisplayCart)

		}

		r.GET("/checkout",handlers.CheckOut)

	}

	return r

}
