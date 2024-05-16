package main

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"Tufind-Backend/routes"

	"github.com/gin-gonic/gin"

	// cors "github.com/rs/cors/wrapper/gin"
	"log"
	"net/http"
)

func main() {
	//initialize database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Tutor{})
	database.DB.AutoMigrate(&models.User{})
	//create gin server
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.TutorRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
