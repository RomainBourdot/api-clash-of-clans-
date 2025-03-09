package login

import (
	"groupie-tracker/models"
	temp "groupie-tracker/templates"
	"net/http"
	"strings"
)

var IsRegistered bool
var UserConnected string
var Id int

func LoginController(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		IsRegistered  bool
		UserConnected string
		Err           bool
		Message       string
	}

	var data Data
	data.IsRegistered = IsRegistered
	data.UserConnected = UserConnected
	// Récupère le message depuis l'URL
	data.Message = r.URL.Query().Get("message")

	if r.FormValue("disconnect") == "true" {
		IsRegistered = false
		UserConnected = ""
		Id = 0
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.FormValue("error") != "" {
		data.Err = true
		temp.Temp.ExecuteTemplate(w, "login", data)
		return
	}

	temp.Temp.ExecuteTemplate(w, "login", data)
}

func LoginTraitement(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Redirect(w, r, "/login?error=champvide", http.StatusSeeOther)
		return
	}

	users := models.JsonRead()

	for _, elem := range users {
		if strings.EqualFold(elem.Username, username) {
			if elem.Password == password {
				IsRegistered = true
				UserConnected = username
				Id = elem.Id
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/login?error=mdp", http.StatusSeeOther)
				return
			}
		}
	}

}
