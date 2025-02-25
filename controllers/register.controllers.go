package controllers

import (
	"fmt"
	"groupie-tracker/models"
	temp "groupie-tracker/templates"
	"net/http"
)

// RegisterController gère l'affichage et le traitement de l'inscription
func RegisterController(w http.ResponseWriter, r *http.Request) {
	// Affichage du formulaire d'inscription en GET
	if r.Method == http.MethodGet {
		temp.Temp.ExecuteTemplate(w, "register", nil)
		return
	}

	// Traitement du formulaire en POST
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		// Vérification des champs
		if username == "" || password == "" || confirmPassword == "" {
			http.Redirect(w, r, "/register?error=champvide", http.StatusSeeOther)
			return
		}
		if password != confirmPassword {
			http.Redirect(w, r, "/register?error=mdp", http.StatusSeeOther)
			return
		}

		// Création d'un nouvel utilisateur
		newUser := models.User{
			Username: username,
			Password: password,
			// Les favoris sont initialisés vides par défaut
		}
		err := models.JsonWrite(newUser)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("/register?error=%s", err.Error()), http.StatusSeeOther)
			return
		}

		// Redirection vers la page de connexion après inscription réussie
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Pour toute autre méthode HTTP, renvoyer une erreur
	http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
}
