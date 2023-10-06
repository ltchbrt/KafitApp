package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/KafitApp/models"
)

func CreateSprint(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Sprint{}
	product1 := []models.Sprint{}
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)

	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")

	db.Where("date", date).Find(&product1)

	 result := len(product1)

	 if result == 0{
		product.Sprint = weight
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)

	 }else{
		db.Where("date", date).Find(&product)
		product.Sprint = weight
		product.UserID = user.Value
		db.Save(&product)
	 }
	
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetSprint(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Sprint{}
	db.Preload("User").Last(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}



