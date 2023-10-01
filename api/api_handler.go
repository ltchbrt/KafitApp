package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "user") {
		CreateUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_user") {
		GetUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "bmi") {
		CreateBMI(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_bmi") {
		GetBMI(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "flex") {
		CreateFlex(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_flex") {
		GetFlex(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "basic") {
		CreateStrength(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_basic") {
		GetStrength(w, r)
		return
	}


	
}
