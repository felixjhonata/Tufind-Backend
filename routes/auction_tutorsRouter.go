package routes

import (
	controller "Tufind-Backend/controllers"
	middlewares "Tufind-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func Auction_TutorRoutes(incomingRoutes *gin.Engine) {
	protected := incomingRoutes.Group("/api/protected").Use(middlewares.Authentication())
	{
		protected.GET("/auctiontutors/:auction_id", controller.GetAuctionTutors)
	}

}
