package register

import (
	"groupie-tracker/login"
	"groupie-tracker/models"
	temp "groupie-tracker/templates"
	"net/http"
	"strings"
)

type Data struct {
	IsRegistered  bool
	UserConnected string
	Err           bool
}

func RegisterController(w http.ResponseWriter, r *http.Request) {

	var data Data

	data.IsRegistered = login.IsRegistered
	data.UserConnected = login.UserConnected

	if r.FormValue("exist") == "username" {
		data.Err = true
		temp.Temp.ExecuteTemplate(w, "register", data)
		return
	}

	temp.Temp.ExecuteTemplate(w, "register", nil)
}

func RegisterTraitement(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	users := models.JsonRead()

	for _, elem := range users {
		if strings.EqualFold(elem.Username, username) {
			http.Redirect(w, r, "/register?exist=username", http.StatusSeeOther)
			return
		}
	}

	newUser := models.User{Username: username, Password: password}

	models.JsonWrite(newUser)

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
