package models

type Tutor struct {
	ID    uint `gorm:"primaryKey"`
	Price uint `json:"price"`
}
