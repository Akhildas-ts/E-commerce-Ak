package routes

import (
	"ak/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	r.POST("/signup", handlers.Signup)
	// r.POST("/login",handlers.LoginWithPassword)

}
