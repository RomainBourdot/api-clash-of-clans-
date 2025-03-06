package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func ListClans(w http.ResponseWriter, r *http.Request) {

	listClans, err := services.GetClanByQuery("neuille", "", "", "")
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", http.StatusPermanentRedirect, err.Error()), http.StatusPermanentRedirect)
		return
	}
	temp.Temp.ExecuteTemplate(w, "clans", listClans)
}

func DetailsClan(w http.ResponseWriter, r *http.Request) {
	tag := r.FormValue("tag")
	if tag == "" {
		http.Error(w, "Le tag du clan ne peut pas être vide", http.StatusBadRequest)
		return
	}

	clan, err := services.GetClanByTag(tag[1:])
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des données : %s", err.Error()), http.StatusInternalServerError)
		return
	}

	temp.Temp.ExecuteTemplate(w, "details-clans", clan)
}
