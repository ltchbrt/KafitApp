package models

import "gorm.io/gorm"

type PARQ struct {
	ID       uint `gorm:"primaryKey"`
	User		User
	UserID		string
	Q1			string
	Q2			string
	Q3			string
	Q4			string
	Q5			string
	Q6			string
	Q7			string
	Date		string
	Deleted  gorm.DeletedAt
}


