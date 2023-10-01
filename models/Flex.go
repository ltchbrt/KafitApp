package models

import "gorm.io/gorm"

type Flex struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Test		float64
	Sit			float64
	Date		string
	Deleted  gorm.DeletedAt
}


