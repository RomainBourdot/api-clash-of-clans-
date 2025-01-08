package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func ListClans(w http.ResponseWriter, r *http.Request) {
	listClans, err := services.GetClanByTag("test")
	// Vérification d'une erreur dans la réponse de l'api
	if err != nil {
		// Redirection vers la page d'erreur en cas d'erreur dans la réponse
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", http.StatusPermanentRedirect, err.Error()), http.StatusPermanentRedirect)
		return
	}
	fmt.Println(listClans)
	// Chargement et rendue du template "clans" avec les données de l'API
	// Envoie de la réponse au client
	temp.Temp.ExecuteTemplate(w, "clans", listClans)
}
