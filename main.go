package main

import (
	"groupie-tracker/routes"
	temp "groupie-tracker/templates"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	temp.InitTemplates()
	routes.InitServe()
}
