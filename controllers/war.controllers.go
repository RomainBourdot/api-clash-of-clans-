package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

// ListWars affiche la liste des guerres de clans.
func ListWars(w http.ResponseWriter, r *http.Request) {
	wars, err := services.GetWars()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des guerres: %v", err), http.StatusInternalServerError)
		return
	}
	temp.Temp.ExecuteTemplate(w, "wars", wars)
}

// WarDetails affiche les détails d'une guerre de clans.
func WarDetails(w http.ResponseWriter, r *http.Request) {
	warID := r.FormValue("id")
	if warID == "" {
		http.Error(w, "Identifiant de guerre manquant", http.StatusBadRequest)
		return
	}
	war, err := services.GetWarDetails(warID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des détails de la guerre: %v", err), http.StatusInternalServerError)
		return
	}
	temp.Temp.ExecuteTemplate(w, "war-details", war)
}
