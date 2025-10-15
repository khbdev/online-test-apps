package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName       string `gorm:"column:first_name;type:varchar(255)"`
	LastName        string `gorm:"column:last_name;type:varchar(255)"`
	Phone           string `gorm:"column:phone;type:varchar(100)"`
	Email           string `gorm:"column:email;type:varchar(255)"`
	TgUsername      string `gorm:"column:tg_username;type:varchar(255)"`
	Bolimlar        string `gorm:"column:bolimlar;type:json"`
	Savollar        string `gorm:"column:savollar;type:json"`
	Javoblar        string `gorm:"column:javoblar;type:json"`
	TogriJavoblar   int    `gorm:"column:togri_javoblar"`
	NatogriJavoblar int    `gorm:"column:natogri_javoblar"`
}
