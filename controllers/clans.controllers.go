package controllers

import (
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
	"strconv"
)

func ListClans(w http.ResponseWriter, r *http.Request) {

	query := r.FormValue("research")
	if query == "" {
		query = "neuille"
	}

	page := 1
	limit := 10
	if pageStr := r.FormValue("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr := r.FormValue("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	listClans, err := services.GetClanByQuery(query, "", "", "")
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", http.StatusPermanentRedirect, err.Error()), http.StatusPermanentRedirect)
		return
	}

	totalItems := len(listClans.Items)
	totalPages := (totalItems + limit - 1) / limit

	if totalPages == 0 {
		page = 0
	} else {
		if page > totalPages {
			page = totalPages
		}
		if page < 1 {
			page = 1
		}
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit
	if endIndex > totalItems {
		endIndex = totalItems
	}

	var paginatedItems interface{}
	if totalItems > 0 {
		paginatedItems = listClans.Items[startIndex:endIndex]
	} else {
		paginatedItems = []interface{}{}
	}

	var prevPage, nextPage int
	if page > 1 {
		prevPage = page - 1
	}
	if page < totalPages {
		nextPage = page + 1
	}

	data := map[string]interface{}{
		"Items":       paginatedItems,
		"CurrentPage": page,
		"TotalPages":  totalPages,
		"PrevPage":    prevPage,
		"NextPage":    nextPage,
		"Query":       query,
		"Limit":       limit,
	}

	err = temp.Temp.ExecuteTemplate(w, "clans", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du rendu du template : %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func DetailsClan(w http.ResponseWriter, r *http.Request) {
	tag := r.FormValue("tag")
	if tag == "" {
		http.Error(w, "Le tag du clan ne peut pas être vide", http.StatusBadRequest)
		return
	}

	clan, err := services.GetClanByTag(tag[1:])
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des données : %s", err.Error()), http.StatusInternalServerError)
		return
	}

	temp.Temp.ExecuteTemplate(w, "details-clans", clan)
}
