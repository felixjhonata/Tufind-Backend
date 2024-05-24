package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetAuctionTutors(c *gin.Context) {
	var auction_tutors []models.Tutor
	database.DB.Find(&auction_tutors)
	c.JSON(http.StatusOK, gin.H{"tutors": auction_tutors})
}
