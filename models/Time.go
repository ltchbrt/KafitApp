package models

import "gorm.io/gorm"

type TimeD struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Drop		float64
	Date		string
	Deleted  gorm.DeletedAt
}


