package controllers

import (
	"Tufind-Backend/database"
	"Tufind-Backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateBidsInput struct {
	Amount float64 `json:"amount"`
}

func UpdateBid(c *gin.Context) {
	var input UpdateBidsInput
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := UpdateBidAmount(database.DB, id, input.Amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bid updated"})
}
func UpdateBidAmount(db *gorm.DB, bidID string, newAmount float64) error {
	// Retrieve the bid from the database
	var bid models.Bid
	result := db.First(&bid, bidID)
	if result.Error != nil {
		return fmt.Errorf("failed to find bid: %v", result.Error)
	}

	// Update the bid amount
	if newAmount > bid.Amount {
		bid.Amount = newAmount
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
