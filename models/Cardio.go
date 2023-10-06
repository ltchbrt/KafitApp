package models

import "gorm.io/gorm"

type Cardio struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Before		float64
	After		float64
	Date		string
	Deleted  gorm.DeletedAt
}


