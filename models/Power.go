package models

import "gorm.io/gorm"

type SpeedT struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Speed		float64
	Speed1		float64
	Date		string
	Deleted  gorm.DeletedAt
}


