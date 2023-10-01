package models

import "gorm.io/gorm"

type Juggling struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Jug		float64
	Date		string
	Deleted  gorm.DeletedAt
}


