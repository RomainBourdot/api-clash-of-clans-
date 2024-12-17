package controllers

import (
	temp "groupie-tracker/templates"
	"net/http"
)

func AccueilController(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "accueil", nil)
}
