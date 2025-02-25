package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Favorite []FavoriteClan `json:"favorite"`
}

type FavoriteClan struct {
	UserID   string `json:"user_id"`
	Tag      string `json:"tag"`
	Name     string `json:"name"`
	Badge    string `json:"badge"`
	Position int    `json:"position"`
}

func JsonRead() []User {

	var data []User

	file, err := os.ReadFile("data.json")
	if err == nil {
		json.Unmarshal(file, &data)
	} else {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}

	return data
}

func JsonWrite(data User) error {

	dataFileJson := JsonRead()

	for _, elem := range dataFileJson {
		if elem.Username == data.Username {
			return fmt.Errorf("Utilisateur déjà existant")
		}
	}

	if len(dataFileJson) == 0 {
		data.Id = 0
	} else {
		data.Id = dataFileJson[len(dataFileJson)-1].Id + 1
	}

	dataFileJson = append(dataFileJson, data)

	file, err := json.MarshalIndent(dataFileJson, "", " ")
	if err != nil {
		return fmt.Errorf("Erreur lors de la conversion en JSON : %v", err)
	}

	err = os.WriteFile("data.json", file, 0644)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'écriture du fichier : %v", err)
	}

	return nil
}

func JsonWriteList(data []User) error {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("Erreur lors de la conversion en JSON : %v", err)
	}

	err = os.WriteFile("data.json", file, 0644)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'écriture du fichier : %v", err)
	}

	return nil
}

func GetUserById(id int) (User, error) {

	Users := JsonRead()

	for _, elem := range Users {
		if elem.Id == id {
			return elem, nil
		}
	}

	return User{}, fmt.Errorf("Utilisateur non trouvé")
}

func FavoriteAdd(data FavoriteClan, idUser int) error {

	file := JsonRead()

	index := -1
	for i, elem := range file {
		if elem.Id == idUser {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Utilisateur non trouvé")
	}

	for _, elem := range file[index].Favorite {
		if elem.Name == data.Name && elem.Tag == data.Tag {
			return fmt.Errorf("Favoris déjà existant")
		}
	}

	file[index].Favorite = append(file[index].Favorite, data)

	JsonWriteList(file)

	return nil
}

func FavoriteDelete(idUser int, tagFavorite string) error {

	file := JsonRead()

	index := -1
	for i, elem := range file {
		if elem.Id == idUser {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Utilisateur non trouvé")
	}

	indexFavorite := -1
	for i, elem := range file[index].Favorite {
		if elem.Tag == tagFavorite {
			indexFavorite = i
			break
		}
	}

	if indexFavorite == -1 {
		return fmt.Errorf("Favoris non trouvé")
	}

	file[index].Favorite = append(file[index].Favorite[:indexFavorite], file[index].Favorite[indexFavorite+1:]...)

	JsonWriteList(file)

	return nil
}
