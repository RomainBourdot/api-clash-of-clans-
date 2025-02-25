package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"groupie-tracker/services"
	"net/http"
)

// getCurrentUserID est une fonction fictive à adapter selon ton système d'authentification.
func getCurrentUserID(r *http.Request) (string, error) {
	// Exemple : récupérer une valeur depuis une session ou un token JWT
	return "user1", nil // À remplacer par la logique réelle
}

func AddFavoriteController(w http.ResponseWriter, r *http.Request) {
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
		Badge:  badge,
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

func ReorderFavoritesController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	userID, err := getCurrentUserID(r)
	if err != nil || userID == "" {
		http.Error(w, "Utilisateur non authentifié", http.StatusUnauthorized)
		return
	}

	var orderedTags []string
	err = json.NewDecoder(r.Body).Decode(&orderedTags)
	if err != nil {
		http.Error(w, "Format de données invalide", http.StatusBadRequest)
		return
	}

	err = services.UpdateFavoritesOrder(userID, orderedTags)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour de l'ordre: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
