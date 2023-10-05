package models

import "gorm.io/gorm"

type Balance struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Stork		float64
	Stork1		float64
	Date		string
	Deleted  gorm.DeletedAt
}
