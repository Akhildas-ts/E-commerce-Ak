package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// Set the start date to June 1, 2023
	startDate := time.Date(2024, time.March, 20, 0, 0, 0, 0, time.UTC)
	endDate := time.Now()

	for currentDate := startDate; currentDate.Before(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		for j := 0; j < randInt(1, 10); j++ {
			d := strconv.Itoa(currentDate.Day()) + " days ago"
			writeToFile("file.txt", d)
			runCommand("git", "add", ".")
			runCommand("git", "commit", "--date", currentDate.Format("Mon Jan 2 15:04:05 2006 -0700"), "-m", "commit")
			runCommand("git", "push", "-u", "origin", "dev")
		}
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func writeToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(content + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}






// package main

// import (
// 	"ak/config"
// 	"ak/database"
// 	"ak/docs"
// 	routes "ak/router"
// 	"log"

// 	"github.com/gin-gonic/gin"

// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// //	@title			Ak eCommerce API
// //	@version		1.0
// //	@description	API for ecommerce website
// //	@securityDefinitions.apiKey	Bearer
// //	@in							header
// //	@name						Authorization
// //	@license.name	Apache 2.0
// //	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// //	@host			localhost:8080
// //	@BasePath		/
// //
// // @schemes	http
// func main() {

// 	docs.SwaggerInfo.Title = "Ecommerce_site"
// 	docs.SwaggerInfo.Description = "Ecommerce shirt selling application suing Golang"
// 	docs.SwaggerInfo.Version = "1.0"
// 	docs.SwaggerInfo.Host = "localhost:8080"
// 	docs.SwaggerInfo.BasePath = ""
// 	docs.SwaggerInfo.Schemes = []string{"http"}

// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("error loading the config file")
// 	}
// 	db, err := database.ConnectDatabase(cfg)
// 	if err != nil {
// 		log.Fatalf("Error connecting to the database: %v", err)
// 	}

// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	routes.UserRoutes(router.Group("/"), db)
// 	routes.AdminRoutes(router.Group("/admin"), db)

// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	err = router.Run("localhost:8080")
// 	if err != nil {

// 		log.Fatalf("Local host Errors %v", err)

// 	}

// }
