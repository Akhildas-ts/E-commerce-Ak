package routes

import (
	"ak/handlers"
	"ak/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/adminlogin", handlers.AdminLogin)
	// r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBord)

	// r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBord)

	// r.Use(middleware.AuthorizationMiddleware())
	// {
	// 	r.GET("/dashboard", handlers.DashBord)

	// }

	r.Use(middleware.AuthorizationMiddleware())
	{
		r.GET("/dashboard", handlers.DashBord)
	}

	return r

}
