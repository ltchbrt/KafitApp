package models

import "gorm.io/gorm"

type Teacher struct {
	ID       uint `gorm:"primaryKey"`
	User     User
	UserID	 uint
	Section	string
	Grade  string
	Deleted  gorm.DeletedAt
}


