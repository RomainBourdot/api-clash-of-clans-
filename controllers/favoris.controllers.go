package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/login"
	"groupie-tracker/models"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
	"strconv"
)

func getCurrentUserID(r *http.Request) (string, error) {
	if !login.IsRegistered {
		return "", fmt.Errorf("Utilisateur non authentifié")
	}
	return strconv.Itoa(login.Id), nil
}

func FavoriteController(w http.ResponseWriter, r *http.Request) {
	userID, err := getCurrentUserID(r)
	if err != nil || userID == "" {
		// Redirige l'utilisateur vers la page de connexion avec un message d'information.
		// Le message doit être encodé en URL.
		http.Redirect(w, r, "/login?message=Veuillez+vous+connecter+pour+voir+vos+favoris", http.StatusSeeOther)
		return
	}

	favs, err := services.ListFavorites(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des favoris: %v", err), http.StatusInternalServerError)
		return
	}
	// Rendu du template "favorites" avec les données des favoris
	err = temp.Temp.ExecuteTemplate(w, "favorites", favs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du rendu du template: %v", err), http.StatusInternalServerError)
		return
	}
}

func AddFavoriteController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Vérification de l'authentification
	userID, err := getCurrentUserID(r)
	if err != nil || userID == "" {
		// Redirection vers la page de connexion si non authentifié
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tag := r.FormValue("tag")
	name := r.FormValue("name")
	badge := r.FormValue("badge")
	if tag == "" || name == "" {
		http.Error(w, "Tag ou nom manquant", http.StatusBadRequest)
		return
	}

	fav := models.FavoriteClan{
		UserID: userID,
		Tag:    tag,
		Name:   name,
		BadgeUrl: struct {
			Medium string `json:"medium"`
		}{
			Medium: badge,
		},
	}

	err = services.AddFavorite(fav)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ajout du favori: %v", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func RemoveFavoriteController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	userID, err := getCurrentUserID(r)
	if err != nil || userID == "" {
		http.Error(w, "Utilisateur non authentifié", http.StatusUnauthorized)
		return
	}

	tag := r.FormValue("tag")
	if tag == "" {
		http.Error(w, "Tag manquant", http.StatusBadRequest)
		return
	}

	err = services.RemoveFavorite(userID, tag)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la suppression du favori: %v", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func ListFavoritesController(w http.ResponseWriter, r *http.Request) {
	userID, err := getCurrentUserID(r)
	if err != nil || userID == "" {
		http.Error(w, "Utilisateur non authentifié", http.StatusUnauthorized)
		return
	}

	favs, err := services.ListFavorites(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des favoris: %v", err), http.StatusInternalServerError)
		return
	}

	// Ici, on peut utiliser un template pour afficher les favoris ou renvoyer du JSON.
	// Pour l'exemple, nous renvoyons du JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favs)
}
