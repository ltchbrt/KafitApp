package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/KafitApp/models"
)

func CreateCardio(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Cardio{}
	product1 := []models.Cardio{}
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")

	db.Where("date", date).Find(&product1)

	 result := len(product1)

	 if result == 0{
		product.Before = weight
		product.After = height
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)

	 }else{
		db.Where("date", date).Find(&product)
		product.Before = weight
		product.After = height
		product.UserID = user.Value
		db.Save(&product)
	 }
	
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetCardio(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Cardio{}
	db.Preload("User").Last(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}



