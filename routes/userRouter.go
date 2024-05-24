package routes

import (
	controller "Tufind-Backend/controllers"
	"Tufind-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	api := incomingRoutes.Group("/api")
	public := api.Group("/public")
	{
		public.POST("/users/login", controller.Login)
		public.POST("/users/signup", controller.SignUp)
	}
	protected := api.Group("/protected").Use(middlewares.Authentication())
	{
		protected.GET("/users/profile", controller.Profile)
	}
}
