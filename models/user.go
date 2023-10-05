package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string 
	Password string
	Name     string
	Type     string
	Number	 string 
	Age		string
	Sex		string
	Deleted  gorm.DeletedAt
}

func (u *User) String() string {
	return u.Name
}
