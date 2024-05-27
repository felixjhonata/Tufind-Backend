package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateBidsInput struct {
	Price uint `json:"Price"`
}

type InputBidProof struct {
	Proof string `json:"proof"`
}

func CreateBid(c *gin.Context) {
	var bid models.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bid.BidTime = time.Now()
	//find tutor
	var auctionTutor models.AuctionTutor
	if err := database.DB.Preload("Tutor").First(&auctionTutor, bid.AuctionTutorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AuctionTutor not found"})
		return
	}

	// Validate the bid amount against the tutor price
	fmt.Printf(" %d %d\n", bid.Price, auctionTutor.Tutor.Price)
	if bid.Price <= auctionTutor.Tutor.Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bid amount must be higher than the tutor price"})
		return
	}

	//update tutor
	id := auctionTutor.TutorID
	err := UpdateTutorPrice(database.DB, id, bid.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&bid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Bid",
	})
}

func UpdateBid(c *gin.Context) {
	var input UpdateBidsInput
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := UpdateBidAmount(database.DB, id, input.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bid updated"})
}
func UpdateBidAmount(db *gorm.DB, bidID string, newAmount uint) error {
	// Retrieve the bid from the database
	var bid models.Bid
	result := db.First(&bid, bidID)
	if result.Error != nil {
		return fmt.Errorf("failed to find bid: %v", result.Error)
	}

	// Update the bid amount
	if newAmount > bid.Price {
		bid.Price = newAmount
	} else {
		return fmt.Errorf("bid must be bigger than start price")
	}

	// Save the updated bid back to the database
	result = db.Save(&bid)
	if result.Error != nil {
		return fmt.Errorf("failed to update bid amount: %v", result.Error)
	}

	fmt.Printf("Bid amount updated successfully for bid ID %d\n", bid.ID)
	return nil
}

func GetBids(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction ID"})
		return
	}
	bids, err := GetBidsByUserID(database.DB, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type BidResponse struct {
		models.Bid
		Status string `json:"status"`
	}

	//find tutor

	var response []BidResponse
	for _, bid := range bids {
		response = append(response, BidResponse{
			Bid:    bid,
			Status: bid.DetermineStatus(bid.AuctionTutor.Tutor.Price),
		})
	}

	c.JSON(http.StatusOK, response)

}
func GetBidsByUserID(db *gorm.DB, userID uint) ([]models.Bid, error) {
	var bids []models.Bid
	err := db.Preload("User"). // Preload the User associated with each Bid
					Preload("AuctionTutor.Auction").Preload("AuctionTutor.Tutor"). // Preload the Auction associated with each AuctionTutor
					Where("user_id = ?", userID).
					Find(&bids).Error
	if err != nil {
		return nil, err
	}
	return bids, nil
}

func AddProof(c *gin.Context) {
	var input InputBidProof
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := AddBidProof(database.DB, id, input.Proof)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "proof updated"})

}

func AddBidProof(db *gorm.DB, bidID string, newproof string) error {
	// Retrieve the bid from the database
	var bid models.Bid
	result := db.First(&bid, bidID)
	if result.Error != nil {
		return fmt.Errorf("failed to find bid: %v", result.Error)
	}

	// Update the bid proof
	bid.Proof = newproof
	bid.Paid = true
	// Save the updated bid back to the database
	result = db.Save(&bid)
	if result.Error != nil {
		return fmt.Errorf("failed to update proof: %v", result.Error)
	}

	fmt.Printf("Bid proof updated successfully for bid ID %d\n", bid.ID)
	return nil
}
