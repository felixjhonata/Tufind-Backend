package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
	"time"
)

func GetAuctionTimer(c *gin.Context) {

	var auction models.Auction

	id := c.Param("id")
	if err := database.DB.First(&auction, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	currentTime := time.Now()
	remainingTime := auction.EndTime.Sub(currentTime)
	if remainingTime < 0 {
		remainingTime = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"remaining_time": remainingTime.String(),
	})
}

func CreateAuction(c *gin.Context) {
	var auction models.Auction
	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&auction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}
