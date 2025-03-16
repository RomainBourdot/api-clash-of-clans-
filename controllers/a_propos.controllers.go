package controllers

import (
	"html/template"
	"net/http"
)

func AProposController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/a_propos.page.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "À propos",
		Message: "Bienvenue sur la page À propos de mon site !",
	}

	err = tmpl.ExecuteTemplate(w, "a-propos", data)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
		return
	}
}
