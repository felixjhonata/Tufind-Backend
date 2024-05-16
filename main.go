package main

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"Tufind-Backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//initialize database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Tutor{})
	database.DB.AutoMigrate(&models.User{})
	//create gin server
	router := gin.Default()
	//cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.TutorRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
