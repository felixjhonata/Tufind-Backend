package models

type Tutor struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	NamaTutor string `gorm:"type:varchar(250)" json:"nama_tutor"`
}
