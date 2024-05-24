package main

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"Tufind-Backend/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//initialize database
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&models.Tutor{}, &models.Bid{}, &models.User{}, &models.Auction{}, &models.AuctionTutor{})
	if err != nil {
		log.Fatalln("Failed to migrate User schema:", err)
	}
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

	routes.UserRoutes(router)
	routes.TutorRoutes(router)
	routes.AuctionRoutes(router)
	routes.Auction_TutorRoutes(router)
	routes.BidRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
