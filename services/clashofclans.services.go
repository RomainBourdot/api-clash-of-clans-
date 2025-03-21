package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var _httpClient = http.Client{
	Timeout: 5 * time.Second,
}

var _token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjUwMTVjNDc1LWYxZmYtNDA1MS1hMTlhLWY5YjA2YjcyZDAxMiIsImlhdCI6MTc0MjE0NTkwNCwic3ViIjoiZGV2ZWxvcGVyLzA5YTc2OTEyLTk3MWQtMjZhMy1hNDY3LTA2YTkxMjMyNzI5YiIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjkxLjE2My43Ny42MiJdLCJ0eXBlIjoiY2xpZW50In1dfQ.DThVZ7W8ouheykc1ufqJMJL1AVbReZ8a2ig5r09db1ry8Zou_YsGH9Hn1TDlIPmB8omQuRdCGMzykdkJDYtceg"

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
		Label   []struct {
			Name     string `json:"name"`
			IconUrls struct {
				Medium string `json:"medium"`
			} `json:"iconUrls"`
		} `json:"labels"`
	} `json:"items"`
}

type ErrorClient struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func GetClanByQuery(query, minClanLevel, minMembers, minClanPoints string) (ShearchClan, error) {
	params := url.Values{}
	params.Add("name", query)

	if minClanLevel != "" {
		params.Add("minClanLevel", minClanLevel)
	}
	if minMembers != "" {
		params.Add("minMembers", minMembers)
	}
	if minClanPoints != "" {
		params.Add("minClanPoints", minClanPoints)
	}

	url := fmt.Sprintf("https://api.clashofclans.com/v1/clans?%s", params.Encode())

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de l'initialisation de la requête")
	}

	req.Header.Set("Authorization", _token)
	req.Header.Set("Accept", "application/json")

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de l'envoi de la requête")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var data ErrorClient
		errDecode := json.NewDecoder(res.Body).Decode(&data)
		if errDecode != nil {
			return ShearchClan{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
		}
		fmt.Println(data)
		return ShearchClan{}, fmt.Errorf("Erreur lors de la récupération des clans : Code : %d, Message : %s", res.StatusCode, res.Status)
	}

	var data ShearchClan
	errDecode := json.NewDecoder(res.Body).Decode(&data)
	if errDecode != nil {
		return ShearchClan{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
	}
	return data, nil
}

type DetailsClan struct {
	Tag         string `json:"tag"`
	Name        string `json:"name"`
	TypeClan    string `json:"type"`
	Description string `json:"description"`
	Location    struct {
		Name string `json:"name"`
	} `json:"location"`
	BadgeUrl struct {
		Medium string `json:"medium"`
	} `json:"badgeUrls"`
	ClanLevel             int `json:"clanLevel"`
	ClanPoints            int `json:"clanPoints"`
	ClanBuilderBasePoints int `json:"clanBuilderBasePoints"`
	ClanCapitalPoints     int `json:"clanCapitalPoints"`
	CapitalLeague         struct {
		Name string `json:"name"`
	} `json:"capitalLeague"`
	RequiredTrophies int    `json:"requiredTrophies"`
	WarFrequency     string `json:"warFrequency"`
	WarWinStreak     int    `json:"warWinStreak"`
	WarWins          int    `json:"warWins"`
	WarTies          int    `json:"warTies"`
	WarLosses        int    `json:"warLosses"`
	Warleague        struct {
		Name string `json:"name"`
	} `json:"warLeague"`
	Members     int `json:"members"`
	MembersList []struct {
		Name          string `json:"name"`
		Role          string `json:"role"`
		TownHallLevel int    `json:"townHallLevel"`
		ExpLevel      int    `json:"expLevel"`
		League        struct {
			Name      string `json:"name"`
			iconsUrls struct {
				Medium string `json:"medium"`
			} `json:"iconsUrls"`
		} `json:"league"`
		Trophies            int `json:"trophies"`
		BuilderBaseTrophies int `json:"builderBaseTrophies"`
		Donations           int `json:"donations"`
		DonationsReceived   int `json:"donationsReceived"`
	} `json:"memberList"`
	Labels []struct {
		Name     string `json:"name"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"labels"`
	RequiredBuilderBaseTrophies int `json:"requiredBuilderBaseTrophies"`
	RequieredTownHallLevel      int `json:"requiredTownHallLevel"`
	ClanCapital                 struct {
		CapitalHallLevel int `json:"capitalHallLevel"`
		Districts        []struct {
			Name              string `json:"name"`
			DistrictHallLevel int    `json:"districtHallLevel"`
		} `json:"districts"`
	} `json:"clanCapital"`
}

func GetClanByTag(tag string) (DetailsClan, error) {
	url := fmt.Sprintf("https://api.clashofclans.com/v1/clans/%%23%s", tag)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return DetailsClan{}, fmt.Errorf("Erreur lors de l'initialisation de la réquête")
	}

	req.Header.Set("Authorization", _token)
	req.Header.Set("Accept", "application/json")

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return DetailsClan{}, fmt.Errorf("Erreur lors de l'envois de la réquête")
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	if res.StatusCode != http.StatusOK {
		var data ErrorClient
		errDecode := json.NewDecoder(res.Body).Decode(&data)
		if errDecode != nil {
			return DetailsClan{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
		}
		fmt.Println(data)
		return DetailsClan{}, fmt.Errorf("Erreur lors de la récupération des clans : \n Code : %d\n Message : %s", res.StatusCode, res.Status)
	}

	var data DetailsClan
	errDecode := json.NewDecoder(res.Body).Decode(&data)
	if errDecode != nil {
		return DetailsClan{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
	}
	return data, nil
}

type WarItem struct {
	Result           string `json:"result"`
	EndTime          string `json:"endTime"`
	TeamSize         int    `json:"teamSize"`
	AttacksPerMember int    `json:"attacksPerMember"`
	BattleModifier   string `json:"battleModifier,omitempty"`
	Clan             struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		BadgeUrls struct {
			Medium string `json:"medium"`
		} `json:"badgeUrls"`
		ClanLevel             int     `json:"clanLevel"`
		Attacks               int     `json:"attacks"`
		Stars                 int     `json:"stars"`
		DestructionPercentage float64 `json:"destructionPercentage"`
		ExpEarned             int     `json:"expEarned"`
	} `json:"clan"`
	Opponent struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		BadgeUrls struct {
			Medium string `json:"medium"`
		} `json:"badgeUrls"`
		ClanLevel             int     `json:"clanLevel"`
		Attacks               int     `json:"attacks"`
		Stars                 int     `json:"stars"`
		DestructionPercentage float64 `json:"destructionPercentage"`
	} `json:"opponent"`
}

type ClanWars struct {
	Items []WarItem `json:"items"`
}

func GetWarsByClanTag(tag string) (ClanWars, error) {

	cleanedTag := strings.TrimPrefix(tag, "#")

	url := fmt.Sprintf("https://api.clashofclans.com/v1/clans/%%23%s/warlog?limit=5", cleanedTag)
	fmt.Println(url)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return ClanWars{}, fmt.Errorf("Erreur lors de l'initialisation de la requête")
	}

	req.Header.Set("Authorization", _token)
	req.Header.Set("Accept", "application/json")

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return ClanWars{}, fmt.Errorf("Erreur lors de l'envoi de la requête")
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	if res.StatusCode != http.StatusOK {
		var data ErrorClient
		errDecode := json.NewDecoder(res.Body).Decode(&data)
		if errDecode != nil {
			return ClanWars{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
		}
		fmt.Println(data)
		return ClanWars{}, fmt.Errorf("Erreur lors de la récupération des clans : \n Code : %d\n Message : %s", res.StatusCode, res.Status)
	}

	var data ClanWars
	errDecode := json.NewDecoder(res.Body).Decode(&data)
	if errDecode != nil {
		return ClanWars{}, fmt.Errorf("Erreur lors de la lecture de la réponse de l'API : %s", errDecode)
	}
	return data, nil
}
