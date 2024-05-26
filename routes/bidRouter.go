package routes

import (
	controller "Tufind-Backend/controllers"
	middlewares "Tufind-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func BidRoutes(incomingRoutes *gin.Engine) {
	protected := incomingRoutes.Group("/api/protected").Use(middlewares.Authentication())
	{
		protected.GET("/bid/:auction_id", controller.GetBids)
		protected.PUT("/bid/:id", controller.UpdateBid)
		protected.PUT("/bid/pay/:id", controller.AddProof)
		protected.POST("/bid", controller.CreateBid)
	}

}
