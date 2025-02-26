package controllers

import (
	"groupie-tracker/login"
	temp "groupie-tracker/templates"
	"net/http"
)

func AccueilController(w http.ResponseWriter, r *http.Request) {
	var data = login.IsRegistered

	temp.Temp.ExecuteTemplate(w, "accueil", data)
}
