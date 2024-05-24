package routes

import (
	controller "Tufind-Backend/controllers"
	"Tufind-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func BidRoutes(incomingRoutes *gin.Engine) {
	protected := incomingRoutes.Group("/api/protected").Use(middlewares.Authentication())
	{
		protected.PUT("/bid/:id", controller.UpdateBid)
	}

}
