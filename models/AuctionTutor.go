package models

type AuctionTutor struct {
	ID        uint    `gorm:"primaryKey"`
	AuctionID uint    `json:"auction_id"`
	Auction   Auction `gorm:"foreignKey:AuctionID" json:"auction"`
	TutorID   uint    `json:"tutor_id"`
	Tutor     Tutor   `gorm:"foreignKey:TutorID" json:"tutor"`
}
