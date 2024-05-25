package routes

import (
	controller "Tufind-Backend/controllers"
	"Tufind-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func AuctionRoutes(incomingRoutes *gin.Engine) {
	protected := incomingRoutes.Group("/api/protected").Use(middlewares.Authentication())
	{
		protected.GET("/auction/timer/:id", controller.GetAuctionTimer)
		protected.POST("/auction", controller.CreateAuction)
		protected.GET("/auction", controller.GetAuction)
	}

}
