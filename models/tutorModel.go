package models

type Tutor struct {
	ID             uint     `gorm:"primaryKey"`
	University     string   `json:"university"`
	Price          uint     `json:"price"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	Email          string   `gorm:"unique" json:"email"`
	KPU            *float64 `json:"kemampuan_penalaran_umum,omitempty"`
	KMBM           *float64 `json:"kemampuan_memahami_bacaan_dan_menulis,omitempty"`
	PPM            *float64 `json:"pengetahuan_dan_pemahaman_umum,omitempty"`
	PK             *float64 `json:"pengetahuan_kuantitatif,omitempty"`
	MathGrade      *float64 `json:"math,omitempty"`
	PhysicsGrade   *float64 `json:"physics,omitempty"`
	BiologyGrade   *float64 `json:"biology,omitempty"`
	ChemistryGrade *float64 `json:"chemistry,omitempty"`
	HistoryGrade   *float64 `json:"history,omitempty"`
	GeographyGrade *float64 `json:"geography,omitempty"`
	SosiologyGrade *float64 `json:"sosiology,omitempty"`
	EconomyGrade   *float64 `json:"economy,omitempty"`
}
