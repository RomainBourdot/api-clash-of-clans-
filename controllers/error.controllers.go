package controllers

import (
	temp "groupie-tracker/templates"
	"net/http"
)

type PageError struct {
	Code    string
	Message string
}

func ErrorController(w http.ResponseWriter, r *http.Request) {
	data := PageError{r.FormValue("code"), r.FormValue("message")}
	temp.Temp.ExecuteTemplate(w, "error", data)
}
