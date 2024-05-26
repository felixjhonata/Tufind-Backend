package routes

import (
	controller "Tufind-Backend/controllers"
	"Tufind-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func TutorRoutes(incomingRoutes *gin.Engine) {

	protected := incomingRoutes.Group("/api/protected").Use(middlewares.Authentication())
	{
		protected.GET("/tutors", controller.GetTutors())
		protected.GET("/tutor/:id", controller.GetTutor())
		protected.POST("/tutor", controller.CreateTutor())
		protected.PATCH("/tutor/:id", controller.UpdateTutor())
	}

}
