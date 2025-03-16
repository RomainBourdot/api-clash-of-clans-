package controllers

import (
	"html/template"
	"net/http"
)

// AProposController gère l'affichage de la page "À propos".
func AProposController(w http.ResponseWriter, r *http.Request) {
	// On vérifie que la méthode est GET (si nécessaire)
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Chargement du template HTML (adapter le chemin si nécessaire)
	tmpl, err := template.ParseFiles("templates/a_propos.page.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	// Exemple de données à transmettre au template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "À propos",
		Message: "Bienvenue sur la page À propos de mon site !",
	}

	// Exécution du template avec les données
	err = tmpl.ExecuteTemplate(w, "a-propos", data)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
		return
	}
}
