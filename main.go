package main

import (
	"ak/config"
	"ak/database"
	routes "ak/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
	}
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	routes.UserRoutes(router.Group("/"), db)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Local host Error %v", err)

	}

}
