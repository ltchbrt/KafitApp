package models

import "gorm.io/gorm"

type BMI struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Weight		float64
	Height		float64
	Result		float64
	Date		string
	Deleted  gorm.DeletedAt
}


