package models

import "gorm.io/gorm"

type TimeD struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Drop		float64
	Drop1		float64
	Drop2		float64
	Date		string
	Deleted  gorm.DeletedAt
}


