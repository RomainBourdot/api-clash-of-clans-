package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

type searchdata struct {
	Name services.ShearchClan
	/*Pseudo services.ClanMembers*/
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

/*func MembersSearch(w http.ResponseWriter, r *http.Request) {
	// Récupère l'entrée utilisateur depuis la requête
	input := r.FormValue("research")
	if input == "" {
		http.Error(w, "La recherche ne peut pas être vide", http.StatusBadRequest)
		return
	}

	// Appelle le service pour récupérer les données des clans
	research, err := services.GetMembersClan(input) // Rendre l'appel dynamique
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des données : %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var data services.ClanMembers // Supposé comme étant le bon type
	for _, v := range research.Items {
		// Comparaison avec l'entrée utilisateur
		if v.Name == input {
			data.Items = append(data.Items, v)
		}
	}

	// Vérifie si aucun résultat n'a été trouvé
	if len(data.Items) == 0 {
		http.Error(w, "Aucun pseudo correspondant trouvé", http.StatusNotFound)
		return
	}

	// Exécute le template avec les données trouvées
	err = temp.Temp.ExecuteTemplate(w, "research", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du rendu du template : %s", err.Error()), http.StatusInternalServerError)
		return
	}
}*/
