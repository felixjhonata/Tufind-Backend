package routes

import (
	controller "Tufind-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func TutorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tutors", controller.GetTutors())
	incomingRoutes.GET("/tutors/:id", controller.GetTutor())
	incomingRoutes.POST("/tutors", controller.CreateTutor())
	incomingRoutes.PATCH("/tutors/:id", controller.UpdateTutor())
}
