package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func ListClans(w http.ResponseWriter, r *http.Request) {

	listClans, err := services.GetClanByQuery("neuille")
	// Vérification d'une erreur dans la réponse de l'api
	if err != nil {
		// Redirection vers la page d'erreur en cas d'erreur dans la réponse
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", http.StatusPermanentRedirect, err.Error()), http.StatusPermanentRedirect)
		return
	}
	// Chargement et rendue du template "clans" avec les données de l'API
	// Envoie de la réponse au client
	temp.Temp.ExecuteTemplate(w, "clans", listClans)
}

func DetailsClan(w http.ResponseWriter, r *http.Request) {
	// Récupère le tag du clan depuis la requête
	tag := r.FormValue("tag")
	if tag == "" {
		http.Error(w, "Le tag du clan ne peut pas être vide", http.StatusBadRequest)
		return
	}

	// Appelle le service pour récupérer les données du clan
	clan, err := services.GetClanByTag(tag[1:])
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des données : %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Chargement et rendue du template "clan" avec les données de l'API
	// Envoie de la réponse au client
	temp.Temp.ExecuteTemplate(w, "details-clans", clan)
}
