package models

import (
	"time"
)

type Auction struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	ItemName    string    `json:"item_name"`
	Description string    `json:"description"`
	StartPrice  float64   `json:"start_price"`
	EndTime     time.Time `json:"end_time"`
	UserID      uint      // Foreign key to User
	User        User      `gorm:"foreignKey:UserID"`
}

// db.Model(Email{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
//	type Bid struct {
//		ID        uint      `gorm:"primaryKey" json:"id"`
//		AuctionID uint      `json:"auction_id"`
//		Auction   Auction   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"auction"`
//		UserID    uint      `json:"user_id"`
//		User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
//		Bidder    string    `json:"bidder"`
//		Amount    float64   `json:"amount"`
//		BidTime   time.Time `json:"bid_time"`
//	}
// type Auction struct {
// 	ID      uint      `gorm:"primaryKey" json:"id"`
// 	UserID  uint      `json:"user_id"`
// 	User    User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
// 	EndTime time.Time `json:"end_time"`
// 	Bid     []Bid     `json:"bid" gorm:"foreignKey:BidID"`
// }
