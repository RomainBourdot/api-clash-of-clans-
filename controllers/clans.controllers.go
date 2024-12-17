package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func pageListClans(w http.ResponseWriter, r *http.Request) {
	listClans, listClansCode, listClansErr := services.GetClan()
	// Vérification d'une erreur dans la réponse de l'api
	if listClansErr != nil {
		// Redirection vers la page d'erreur en cas d'erreur dans la réponse
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=Erreur lors de la récupération des clans", listClansCode), http.StatusPermanentRedirect)
		return
	}
	// Chargement et rendue du template "albums" avec les données de l'API
	// Envoie de la réponse au client
	temp.Temp.ExecuteTemplate(w, "clans", listClans)
}
