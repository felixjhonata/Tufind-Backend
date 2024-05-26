package models

import (
	"time"
)

const (
	BidStatusInvalid = "invalid"
	BidStatusValid   = "valid"
	BidStatusNotPaid = "not_paid"
	BidStatusDone    = "done"
)

type Bid struct {
	ID             int          `gorm:"primarykey"`
	UserID         uint         `json:"user_id"`
	User           User         `gorm:"foreignKey:UserID" json:"user"`
	AuctionTutorID uint         `json:"auctiontutor_id"`
	AuctionTutor   AuctionTutor `gorm:"foreignKey:AuctionTutorID" json:"auctiontutor"`
	Price          uint         `json:"price"`
	Session        int          `json:"session"`
	BidTime        time.Time    `json:"bid_time"`
	Paid           bool         `json:"paid"`
	Proof          string       `json:"proof,omitempty"`
}

func (b *Bid) DetermineStatus(price uint) string {
	if b.Price == price {
		return BidStatusValid
	}
	if b.AuctionTutor.Auction.IsDone() {
		if b.Paid {
			return BidStatusDone
		}
		return BidStatusNotPaid
	}
	return BidStatusInvalid
}
