package models

import (
	"time"
)

// type Bid struct {
// 	ID        uint      `gorm:"primaryKey" json:"id"`
// 	AuctionID uint      `json:"auction_id"`
// 	Auction   Auction   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"auction"`
// 	UserID    uint      `json:"user_id"`
// 	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
// 	Bidder    string    `json:"bidder"`
// 	Amount    float64   `json:"amount"`
// 	BidTime   time.Time `json:"bid_time"`
// }

type Bid struct {
	ID        int       `json:"id" gorm:"primarykey"`
	AuctionID int       `json:"auction_id"`
	Bidder    string    `json:"bidder"`
	Amount    float64   `json:"amount"`
	BidTime   time.Time `json:"bid_time"`
}
