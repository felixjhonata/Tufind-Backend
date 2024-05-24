package models

import (
	"time"

	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	ID        int       `json:"id" gorm:"primarykey"`
	AuctionID uint      `json:"auction_id"`
	Auction   Auction   `gorm:"foreignKey:AuctionID" json:"auctions"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Amount    float64   `json:"amount"`
	Session   int       `json:"session"`
	BidTime   time.Time `json:"bid_time"`
}
