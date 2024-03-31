package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:255;not null"`
}
