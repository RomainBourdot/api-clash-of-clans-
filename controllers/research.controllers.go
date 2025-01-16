package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

type searchdata struct {
	Name services.ShearchClan
}

func ResearchData(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("research")
	if query == "" {
		http.Error(w, "La recherche ne peut pas être vide", http.StatusBadRequest)
		return
	}

	listClans, err := services.GetClanByQuery(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des données : %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if len(listClans.Items) == 0 {
		http.Error(w, "Aucun clan correspondant trouvé", http.StatusNotFound)
		return
	}

	err = temp.Temp.ExecuteTemplate(w, "research-clans", listClans)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du rendu du template : %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
