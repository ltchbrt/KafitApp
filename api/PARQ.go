package api

import (
	"net/http"
	"time"

	"github.com/dafalo/KafitApp/models"
)

func CreatePARQ(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.PARQ{}
	product1 := []models.PARQ{}
	Q1 := r.FormValue("q1")
	Q2 := r.FormValue("q2")
	Q3 := r.FormValue("q3")
	Q4 := r.FormValue("q4")
	Q5 := r.FormValue("q5")
	Q6 := r.FormValue("q6")
	Q7 := r.FormValue("q7")
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")

	db.Where("user_id", user.Value).Find(&product1)

	 result := len(product1)

	 if result == 0{
		product.Q1 = Q1
		product.Q2 = Q2
		product.Q3 = Q3
		product.Q4 = Q4
		product.Q5 = Q5
		product.Q6 = Q6
		product.Q7 = Q7
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)

	 }else{
		db.Where("user_id", user.Value).Find(&product)
		product.Q1 = Q1
		product.Q2 = Q2
		product.Q3 = Q3
		product.Q4 = Q4
		product.Q5 = Q5
		product.Q6 = Q6
		product.Q7 = Q7
		product.UserID = user.Value
		product.Date = date
		db.Save(&product)
	 }
	
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetPARQ(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	user, _ := r.Cookie("id")

	item := []models.PARQ{}
	db.Preload("User").Where("user_id", user.Value).Last(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}



