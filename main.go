package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dafalo/KafitApp/api"
	"github.com/dafalo/KafitApp/models"
	"github.com/dafalo/KafitApp/views"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":2027"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("kafit")
	MigrateDB()
	CreateDefaultUser()
	Handlers()

	http.ListenAndServe(Port, nil)
}



func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", views.LoginHandler)
	http.HandleFunc("/overview", views.OverviewHandler)
	http.HandleFunc("/PFT", views.PFTHandler)
	http.HandleFunc("/About", views.AboutHandler)
	http.HandleFunc("/Observation", views.ObservationHandler)
	http.HandleFunc("/Materials", views.MaterialsHandler)
	http.HandleFunc("/Objectives", views.ObjectivesHandler)
	http.HandleFunc("/BMI", views.BMIHandler)
	http.HandleFunc("/Flexibility", views.FlexibilityHandler)
	http.HandleFunc("/Cardio", views.CardioHandler)
	http.HandleFunc("/Strength", views.StrengthHandler)
	http.HandleFunc("/Speed", views.SpeedHandler)
	http.HandleFunc("/Power", views.PowerHandler)
	http.HandleFunc("/Agility", views.AgilityHandler)
	http.HandleFunc("/Time", views.TimeHandler)
	http.HandleFunc("/Coordination", views.CoordinationHandler)
	http.HandleFunc("/Balance", views.BalanceHandler)
	http.HandleFunc("/BMI_Weight", views.BMI_WHandler)
	http.HandleFunc("/BMI_Height", views.BMI_HHandler)
	http.HandleFunc("/Flexibility_Zipper", views.Flexibility_ZHandler)
	http.HandleFunc("/Flexibility_Sit", views.Flexibility_SHandler)
	http.HandleFunc("/Cardio_Test", views.CardioTestHandler)
	http.HandleFunc("/Strength_Push", views.Strength_PHandler)
	http.HandleFunc("/Strength_Plank", views.Strength_PLHandler)
	http.HandleFunc("/Speed_Meter", views.Speed_MHandler)
	http.HandleFunc("/Power_Long", views.Power_LHandler)
	http.HandleFunc("/Agility_Test", views.Agility_TestHandler)
	http.HandleFunc("/Time_Drop", views.TimeDropHandler)
	http.HandleFunc("/Coordination_Jug", views.Coordination_JHandler)
	http.HandleFunc("/Balance_Page", views.BalanceCopyHandler)
	http.HandleFunc("/Resources", views.ResourcesHandler)
	http.HandleFunc("/api/", api.APIHandler)
	http.HandleFunc("/logout", views.LogOutHandler)





}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	user := models.User{}
	


	db := GormDB()
	db.AutoMigrate(&user)
}



func CreateDefaultUser() {

	db := GormDB()

	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username: "admin",
			Password: hashPassword("admin"),
			Name:     "Software Developer",
			Type:     "Administrator",
			
			
		},

		{
			Username: "user",
			Password: hashPassword("user"),
			Name:     "Software Developer",
			Type:     "User",
		
		},

		
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}

		if !isExisting {
			fmt.Println("Create Default User")
			db.Save(&defaultUser[i])
		}
	}

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/kafit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

