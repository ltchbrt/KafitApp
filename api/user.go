package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/KafitApp/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	number := r.FormValue("number")
	sex := r.FormValue("sex")
	age := r.FormValue("age")
	section := r.FormValue("section")
	grade := r.FormValue("grade")
	teacher := r.FormValue("teacher")
	section1 := r.FormValue("section1")
	process := r.FormValue("process")


	if process == "teacher"{

		product.Name = name
		product.Username = username
		product.Password = hashPassword(password)
		product.Type = "Teacher"
		product.Number = number
		product.Sex = sex
		product.Age = age
	
		db.Save(&product)

		product1 := models.Teacher{}
		product1.UserID = product.ID
		product1.Section = section
		product1.Grade = grade

		db.Save(&product1)

	}else{
		product.Name = name
		product.Username = username
		product.Password = hashPassword(password)
		product.Type = "User"
		product.Number = number
		product.Sex = sex
		product.Age = age
		product.Teacher = teacher
		product.Section = section1
		db.Save(&product)
	}

	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func getDOB(year, month, day int) time.Time {
    dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return dob
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.User{}
	db.Find(&item)

	item1 := []models.Teacher{}
	db.Preload("User").Find(&item1)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"item1":   item1,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	number := r.FormValue("number")
	
	db.Where("id", id).Find(&product)

	product.Name = name
	product.Username = username
	product.Number = number

	if (password == ""){

	}else{
		product.Password = hashPassword(password)
	}
	db.Save(&product)

}


func EditTeacher(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	uid, _ := strconv.Atoi(r.FormValue("uid"))
	product := models.User{}
	product1 := models.Teacher{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	number := r.FormValue("number")
	sex := r.FormValue("sex")
	section := r.FormValue("section")
	grade := r.FormValue("grade")
	
	db.Where("id", uid).Find(&product)

	product.Name = name
	product.Username = username
	product.Number = number
	product.Sex = sex

	db.Save(&product)

	db.Where("id", id).Find(&product1)
	product1.UserID = product.ID
	product1.Section = section
	product1.Grade = grade

	db.Save(&product1)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.User{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func Password(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	password := r.FormValue("password")
	item := models.User{}
	user, _ := r.Cookie("id")
	db.Where("id", user.Value).Find(&item)

	item.Password = hashPassword(password)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


func DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	uid, _ := strconv.Atoi(r.FormValue("uid"))
	item := models.User{}
	db.Where("id", uid).Statement.Delete(&item)

	item1 := models.Teacher{}
	db.Where("id", id).Statement.Delete(&item1)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}
