package views

import (
	"net/http"
	"text/template"
)

func TLoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/teacher_login.html"))
	data := map[string]interface{}{}
	data["Title"] = "Main | POS_SYSTEM"
	tmpl.Execute(w, data)
}
