package models

import "gorm.io/gorm"

type Agility struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Test		float64
	Date		string
	Deleted  gorm.DeletedAt
}


