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

		if strings.HasPrefix(r.URL.Path, "edit_teacher") {
		EditTeacher(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_teacher") {
		DeleteTeacher(w, r)
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
	if strings.HasPrefix(r.URL.Path, "speed") {
		CreateSpeed(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_speed") {
		GetSpeed(w, r)
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
	if strings.HasPrefix(r.URL.Path, "agility") {
		CreateAgility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_agility") {
		GetAgility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "time") {
		CreateTime(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_time") {
		GetTime(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "jug") {
		CreateCoordination(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_jug") {
		GetCoordination(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "balance") {
		CreateBalance(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_balance") {
		GetBalance(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "cardio") {
		CreateCardio(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_cardio") {
		GetCardio(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "sprint") {
		CreateSprint(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_sprint") {
		GetSprint(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "parq") {
		CreatePARQ(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_parq") {
		GetPARQ(w, r)
		return
	}
}
