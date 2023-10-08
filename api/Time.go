package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/KafitApp/models"
)

func CreateTime(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.TimeD{}
	product1 := []models.TimeD{}
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	weight1, _ := strconv.ParseFloat(r.FormValue("weight1"), 64)
	weight2, _ := strconv.ParseFloat(r.FormValue("weight2"), 64)
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")

	db.Where("date", date).Find(&product1)

	 result := len(product1)

	 if result == 0{
		product.Drop = weight
		product.Drop1 = weight1
		product.Drop2 = weight2
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)

	 }else{
		db.Where("date", date).Find(&product)
		product.Drop = weight
		product.Drop1 = weight1
		product.Drop2 = weight2
		product.UserID = user.Value
		db.Save(&product)
	 }
	
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetTime(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	user, _ := r.Cookie("id")
	item := []models.TimeD{}
	db.Preload("User").Where("user_id", user.Value).Last(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}



