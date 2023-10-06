package models

import "gorm.io/gorm"

type Sprint struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Sprint		float64
	Date		string
	Deleted  gorm.DeletedAt
}


