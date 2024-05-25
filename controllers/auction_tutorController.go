package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAuctionTutors(c *gin.Context) {
	auctionIDStr := c.Param("auction_id")
	auctionID, err := strconv.Atoi(auctionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction ID"})
		return
	}
	auctiontutors, err := GetAuctionTutorsByAuctionID(database.DB, uint(auctionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auctiontutors)

}
func GetAuctionTutorsByAuctionID(db *gorm.DB, auctionID uint) ([]models.AuctionTutor, error) {
	var auctiontutors []models.AuctionTutor
	err := db.Preload("Auction"). // Preload the User associated with each Bid
					Preload("Tutor"). // Preload the Auction associated with each AuctionTutor
					Where("auction_tutors.auction_id = ?", auctionID).
					Find(&auctiontutors).Error
	if err != nil {
		return nil, err
	}
	return auctiontutors, nil
}
