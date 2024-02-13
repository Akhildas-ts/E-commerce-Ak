package main

import (
	"ak/config"
	"ak/database"
	"ak/docs"
	routes "ak/router"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Ak eCommerce API
//	@version		1.0
//	@description	API for ecommerce website
//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			arundas.cloud
//	@BasePath		/
//
// @schemes	http
func main() {


	docs.SwaggerInfo.Title = "Ecommerce_site"
	docs.SwaggerInfo.Description = "Ecommerce shirt selling application suing Golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "arundas.cloud"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
	}
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	routes.UserRoutes(router.Group("/"), db)
	routes.AdminRoutes(router.Group("/admin"), db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("error from ")
		log.Fatalf("Local host Error %v", err)

	}

}
