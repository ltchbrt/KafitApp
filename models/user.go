package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string 
	Password string
	Name     string
	Type     string
	Number	 string 
	Deleted  gorm.DeletedAt
}

func (u *User) String() string {
	return u.Name
}
