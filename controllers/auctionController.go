package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
	"time"
)

type CreateAuctionInput struct {
	ExpectedEndTime time.Time `json:"expected_end_time" binding:"required"`
	TutorIDs        []uint    `json:"tutor_ids" binding:"required"`
}

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
	remainingTime := auction.ExpectedEndTime.Sub(currentTime)
	if remainingTime < 0 {
		remainingTime = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"remaining_time": remainingTime.String(),
	})
}

func CreateAuction(c *gin.Context) {
	// var auction models.Auction
	var input CreateAuctionInput
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load WIB location"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// change all tutor's price to 30
	err1 := database.DB.Model(&models.Tutor{}).Where("true").Update("price", 30)
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	// Create the auction
	auction := models.Auction{ExpectedEndTime: input.ExpectedEndTime.In(wib)}
	database.DB.Create(&auction)

	// Add tutors to the auction
	for _, tutorID := range input.TutorIDs {
		auctionTutor := models.AuctionTutor{AuctionID: auction.ID, TutorID: tutorID}
		database.DB.Create(&auctionTutor)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created Auction",
	})
}

func GetAuction(c *gin.Context) {
	var auction models.Auction
	auction, err := GetLastCreatedAuction(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)

}

func GetLastCreatedAuction(db *gorm.DB) (models.Auction, error) {
	var auction models.Auction
	// Perform the query and check the error
	err := db.Order("id desc").First(&auction).Error
	if err != nil {
		return auction, err
	}
	return auction, nil
}
