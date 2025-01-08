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
var _token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjU2ODE3MmMyLTYxNTYtNDg3YS05ZDg5LWM4ZTI0NDE0NmUyNyIsImlhdCI6MTczNjE1MTQyNiwic3ViIjoiZGV2ZWxvcGVyLzA5YTc2OTEyLTk3MWQtMjZhMy1hNDY3LTA2YTkxMjMyNzI5YiIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjU0Ljg2LjUwLjEzOSJdLCJ0eXBlIjoiY2xpZW50In1dfQ.0sycKTjSSuffRIQ1B0DNqZXzRkLOgmBo4nPFCohjYhljIcq4_i2mFAybwHjgspDKAbqDeOC1Uxck1y8ubXd0zA"

func GetClan(url string) (*http.Response, error) {
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

/* type Clan struct {
	Tag              string   `json:"tag"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Location         Location `json:"location"`
	BadgeUrl         BadgeUrl `json:"badgeUrls"`
	Level            int      `json:"clanLevel"`
	Points           int32    `json:"clanPoints"`
	VersusPoints     int32    `json:"clanVersusPoints"`
	RequiredTrophies int      `json:"requiredTrophies"`
	WarFrequency     string   `json:"warFrequency"`
	WarWinStreak     int      `json:"warWinStreak"`
	WarWins          int      `json:"warWins"`
	WarTies          int      `json:"warTies"`
	WarLosses        int      `json:"warLosses"`
	IsWarLogPublic   bool     `json:"isWarLogPublic"`
	MemberCount      int      `json:"members"`
	Members          []Member `json:"memberList"`
	Labels           []Label  `json:"labels"`
}

// BadgeUrl holds the URL to badge images in various sizes
type BadgeUrl struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// Member holds information about a clan member
type Member struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              string `json:"role"`
	ExpLevel          int    `json:"expLevel"`
	League            League `json:"league"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
	Rank              int    `json:"clanRank"`
	PreviousRank      int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
}


type Location struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsCountry bool   `json:"isCountry"`
}

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"iconUrls"`
}

type Label struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"iconUrls"`
}

type War struct {
	Result   string  `json:"result"`
	Clan     WarClan `json:"clan"`
	Opponent WarClan `json:"opponent"`
}

type WarClan struct {
	Tag                   string      `json:"tag"`
	Name                  string      `json:"name"`
	BadgeUrl              BadgeUrl    `json:"badgeUrls"`
	Level                 int         `json:"clanLevel"`
	Attacks               int         `json:"attacks"`
	Stars                 int         `json:"stars"`
	DestructionPercentage float32     `json:"destructionPercentage"`
	Team                  []WarMember `json:"members"`
}

type WarMember struct {
	Tag      string `json:"tag"`
	Name     string `json:"name"`
	Stars    int    `json:"stars"`
	Attacks  int    `json:"attacks"`
	ExpLevel int    `json:"expLevel"`
} */

type ShearchClan struct {
	Items []struct {
		Tag      string `json:"tag"`
		Name     string `json:"name"`
		TypeClan string `json:"type"`
		BadgeUrl struct {
			Medium string `json:"medium"`
		} `json:"badgeUrls"`
		ClanLevel     int `json:"clanLevel"`
		ClanPoints    int `json:"clanPoints"`
		CapitalLeague struct {
			Name string `json:"name"`
		} `json:"capitalLeague"`
		WarWinStreak int `json:"warWinStreak"`
		WarWins      int `json:"warWins"`
		WarTies      int `json:"warTies"`
		WarLosses    int `json:"warLosses"`
		Warleague    struct {
			Name string `json:"name"`
		} `json:"warLeague"`
		Members int `json:"members"`
	} `json:"items"`
	Label []struct {
		Name     string `json:"name"`
		IconUrls struct {
			Medium string `json:"medium"`
		}
	}
}

func GetClanByTag(name string) (ShearchClan, error) {
	url := fmt.Sprintf("https://api.clashofclans.com/v1/clans?name=%s", name)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de l'initialisation de la réquête")
	}

	req.Header.Set("Authorization", _token)

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de l'envois de la réquête")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ShearchClan{}, fmt.Errorf("Erreur lors de la récupération des clans : \n Code : %d\n Message : %s", res.StatusCode, res.Status)
	}

	var data ShearchClan
	errDecode := json.NewDecoder(res.Body).Decode(&data)
	if errDecode != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
	}
	return data, nil
}
