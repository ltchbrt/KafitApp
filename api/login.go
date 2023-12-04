package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/dafalo/KafitApp/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
   b := make([]byte, length)
   for i := range b {
      b[i] = charset[seededRand.Intn(len(charset))]
   }
   return string(b)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := models.User{}
	report := []models.User{}
	
	db.Where("username = ?", username).Find(&user)
	db.Where("username = ?", username).Find(&report)

	println(username, user.ID)

	db.Where("id", user.ID).Find(&user)

	if CheckPasswordHash(password, user.Password) {
		result := "1"

		newSession := uuid.NewString()

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "session",
			Value: newSession,
		})

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "id",
			Value: fmt.Sprint(user.ID),
		})


		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "tid",
			Value: fmt.Sprint(user.ID),
		})
		data := map[string]interface{}{
			"status":  "ok",
			"results": result,
			"reports": report,
		}
		ReturnJSON(w, r, data)
	} else {
		result := "0"
		data := map[string]interface{}{
			"status":  "error",
			"results": result,
		}
		ReturnJSON(w, r, data)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func GormDB() *gorm.DB {
	dsn := "root:GroupNB2023@tcp(127.0.0.1:3306)/kafit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}


func Email(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	user := models.User{}
	email := r.FormValue("email")
	password := generateRandomString(10)
	
	db.Where("number = ?", email).Find(&user)
	count := "0"
	if user.Number == ""{
		count = "0"
	}else{
		count = "1"
	}

	data := map[string]interface{}{
		"status":  "ok",
		"count": count,
	}
	
	if count == "1" {
		from := "dffalo.amg.pps@gmail.com"
		pass := "hsejfkwaabbkiqrr"
		to := user.Number
	
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Change Password\n\n" +
		   "Good day,\n\n Your New Password to KafitApp is" + " " + password + "\n\n Regards \n\n KafitApp "
	
		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))
	
		if err != nil {
			fmt.Printf("smtp error: %s", err)
			return
		}
		fmt.Println("Successfully sended to " + to)

		fmt.Println("password",password)

		user.Password = hashPassword(password)
		db.Save(&user)
	}

	ReturnJSON(w, r, data)
}




