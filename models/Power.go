package models

import "gorm.io/gorm"

type SpeedT struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Speed		float64
	Date		string
	Deleted  gorm.DeletedAt
}


