package models

import "gorm.io/gorm"

type Strength struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Push		float64
	Basic			float64
	Date		string
	Deleted  gorm.DeletedAt
}


