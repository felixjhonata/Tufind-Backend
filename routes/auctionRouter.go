package routes

import (
	controller "Tufind-Backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuctionRoutes(incomingRoutes *gin.Engine) {
	route := incomingRoutes.Group("/api")
	route.GET("/auction/timer/:id", controller.GetAuctionTimer)
	route.POST("/auction", controller.CreateAuction)
}
