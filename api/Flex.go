package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/KafitApp/models"
)

func CreateFlex(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Flex{}
	product1 := []models.Flex{}
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	left, _ := strconv.ParseFloat(r.FormValue("left"), 64)
	try, _ := strconv.ParseFloat(r.FormValue("try"), 64)
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")

	db.Where("date", date).Find(&product1)

	 result := len(product1)

	 if result == 0{
		product.Test = weight
		product.Left = left
		product.Sit = height
		product.Try = try
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)

	 }else{
		db.Where("date", date).Find(&product)
		product.Test = weight
		product.Left = left
		product.Sit = height
		product.Try = try
		product.UserID = user.Value
		db.Save(&product)
	 }
	
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetFlex(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	user, _ := r.Cookie("id")

	item := []models.Flex{}
	db.Preload("User").Where("user_id", user.Value).Last(&item)

	item1 := []models.Flex{}
	db.Preload("User").Where("user_id", user.Value).Find(&item1)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"item1":   item1,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}



