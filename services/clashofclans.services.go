package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Déclaration du client qui va émettre les requêtes
var _httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// Valeur par défaut pour éviter l'erreur 400 due à un problème de format, type...
var _token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjMzZWZiNzhhLTU0NzItNGUxMC04MDc5LWE4YWM5MDI1MzI3OCIsImlhdCI6MTczNDQyOTkwOCwic3ViIjoiZGV2ZWxvcGVyLzA5YTc2OTEyLTk3MWQtMjZhMy1hNDY3LTA2YTkxMjMyNzI5YiIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjU0Ljg2LjUwLjEzOSJdLCJ0eXBlIjoiY2xpZW50In1dfQ._W7fQdUwAYh50z80EHmLW8tMJHDtckHVNEwkQ2YPrAQWobz5m9S0_menzXZ5Lv2SLbbB_S-4ZKu1gxynDW86-Q"

func requestClashofclansGet(url string) (*http.Response, error) {
	fmt.Printf("valeur token : %s\n", _token)
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return nil, fmt.Errorf("Erreur lors de l'initialisation de la réquête")
	}

	req.Header.Set("Authorization", _token)

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("Erreur lors de l'envois de la réquête")
	}
	return res, nil
}

type ClanData struct {
	ClanName string `json:"name"` // Nom du clan
	ClanTag  string `json:"tag"`  // Tag unique du clan
	ClanInfo struct {
		Level             int    `json:"clanLevel"`             // Niveau du clan
		Points            int    `json:"clanPoints"`            // Points du clan (village principal)
		BuilderBasePoints int    `json:"clanBuilderBasePoints"` // Points de la base du constructeur
		CapitalPoints     int    `json:"clanCapitalPoints"`     // Points de la capitale du clan
		Description       string `json:"description"`           // Description du clan
	} `json:"info"`
	CapitalLeague struct {
		ID   int    `json:"id"`   // ID de la ligue capitale
		Name string `json:"name"` // Nom de la ligue capitale
	} `json:"capitalLeague"`
	Badge struct {
		Urls []struct {
			Small  string `json:"small"`  // URL pour le badge en petite taille
			Medium string `json:"medium"` // URL pour le badge en taille moyenne
			Large  string `json:"large"`  // URL pour le badge en grande taille
		} `json:"badges"`
	} `json:"badgeUrls"`
	WarStats struct {
		WarWins   int `json:"warWins"`   // Nombre de guerres gagnées
		WarLosses int `json:"warLosses"` // Nombre de guerres perdues
		WarTies   int `json:"warTies"`   // Nombre de guerres nulles
	} `json:"warStats"`
	Members []struct {
		Name       string `json:"name"`       // Nom du membre
		Role       string `json:"role"`       // Rôle dans le clan
		ExpLevel   int    `json:"expLevel"`   // Niveau d'expérience
		Trophies   int    `json:"trophies"`   // Nombre de trophées
		LeagueName string `json:"leagueName"` // Nom de la ligue
	} `json:"members"`
}

// Fonction pour récupérer les données du clan
func GetClan(tag, token string) (*ClanData, error) {
	url := fmt.Sprintf("https://api.clashofclans.com/v1/clans/%s", tag)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Ajouter l'en-tête d'authentification
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error code: %d", res.StatusCode)
	}

	var clan ClanData
	if err := json.NewDecoder(res.Body).Decode(&clan); err != nil {
		return nil, err
	}

	return &clan, nil
}
