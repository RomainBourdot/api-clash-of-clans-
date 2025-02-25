package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"groupie-tracker/models"
)

// favoritesFile est le chemin vers le fichier JSON qui stocke les favoris.
var favoritesFile = "./favorites.json"

// favoritesMutex permet d'assurer l'accès concurrent au fichier.
var favoritesMutex sync.Mutex

// FavoritesData représente la structure globale des favoris par utilisateur.
type FavoritesData map[string][]models.FavoriteClan

// loadFavorites lit le fichier JSON et retourne les données.
func loadFavorites() (FavoritesData, error) {
	if _, err := os.Stat(favoritesFile); os.IsNotExist(err) {
		// Si le fichier n'existe pas, retourne une map vide.
		return make(FavoritesData), nil
	}

	data, err := ioutil.ReadFile(favoritesFile)
	if err != nil {
		return nil, err
	}

	var favData FavoritesData
	err = json.Unmarshal(data, &favData)
	if err != nil {
		return nil, err
	}
	return favData, nil
}

// saveFavorites enregistre les données dans le fichier JSON.
func saveFavorites(favData FavoritesData) error {
	data, err := json.MarshalIndent(favData, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(favoritesFile, data, 0644)
}

// AddFavorite ajoute un clan aux favoris pour un utilisateur.
// Il assigne automatiquement la position suivante.
func AddFavorite(fav models.FavoriteClan) error {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favData, err := loadFavorites()
	if err != nil {
		return err
	}

	userFavs := favData[fav.UserID]
	// Détermine la nouvelle position (dernier + 1)
	newPosition := 1
	if len(userFavs) > 0 {
		lastFav := userFavs[len(userFavs)-1]
		newPosition = lastFav.Position + 1
	}
	fav.Position = newPosition

	// Vérifie si le favori existe déjà et le met à jour, sinon l'ajoute.
	updated := false
	for i, existing := range userFavs {
		if existing.Tag == fav.Tag {
			userFavs[i] = fav
			updated = true
			break
		}
	}
	if !updated {
		userFavs = append(userFavs, fav)
	}
	favData[fav.UserID] = userFavs

	return saveFavorites(favData)
}

// RemoveFavorite supprime un favori pour un utilisateur via son tag.
func RemoveFavorite(userID, tag string) error {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favData, err := loadFavorites()
	if err != nil {
		return err
	}

	userFavs, exists := favData[userID]
	if !exists {
		return nil
	}

	newFavs := []models.FavoriteClan{}
	for _, fav := range userFavs {
		if fav.Tag != tag {
			newFavs = append(newFavs, fav)
		}
	}
	favData[userID] = newFavs
	return saveFavorites(favData)
}

// ListFavorites retourne la liste des favoris d'un utilisateur, triés par position.
func ListFavorites(userID string) ([]models.FavoriteClan, error) {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favData, err := loadFavorites()
	if err != nil {
		return nil, err
	}
	return favData[userID], nil
}

// UpdateFavoritesOrder met à jour l'ordre des favoris d'un utilisateur
// en fonction de la liste ordonnée de tags fournie.
func UpdateFavoritesOrder(userID string, orderedTags []string) error {
	favoritesMutex.Lock()
	defer favoritesMutex.Unlock()

	favData, err := loadFavorites()
	if err != nil {
		return err
	}

	userFavs, exists := favData[userID]
	if !exists {
		return nil
	}

	// Crée une map pour retrouver rapidement chaque favori par son tag.
	favMap := make(map[string]models.FavoriteClan)
	for _, fav := range userFavs {
		favMap[fav.Tag] = fav
	}

	newFavs := []models.FavoriteClan{}
	for pos, tag := range orderedTags {
		if fav, ok := favMap[tag]; ok {
			fav.Position = pos + 1
			newFavs = append(newFavs, fav)
		}
	}
	favData[userID] = newFavs
	return saveFavorites(favData)
}
