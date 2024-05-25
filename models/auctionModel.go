package models

import (
	"time"
)

type Auction struct {
	ID              uint      `gorm:"primarykey"`
	ExpectedEndTime time.Time `json:"expected_end_time"`
}

func (a *Auction) IsDone() bool {
	return time.Now().After(a.ExpectedEndTime)
}
