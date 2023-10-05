package models

import "gorm.io/gorm"

type Flex struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Test		float64
	Left		float64
	Sit			float64
	Try			float64
	Date		string
	Deleted  gorm.DeletedAt
}


