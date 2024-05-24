package models

import (
	"time"
)

type Auction struct {
	ID              uint      `gorm:"primarykey"`
	Tutors          []Tutor   `gorm:"many2many:auction_tutors;" json:"tutors"`
	ExpectedEndTime time.Time `json:"expected_end_time"`
}
