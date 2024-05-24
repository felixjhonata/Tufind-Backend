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
		protected.GET("/tutors/:id", controller.GetTutor())
		protected.POST("/tutors", controller.CreateTutor())
		protected.PATCH("/tutors/:id", controller.UpdateTutor())
	}

}
