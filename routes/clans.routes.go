package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func ClansRoutes() {

	http.HandleFunc("/clans", controllers.ListClans)

	http.HandleFunc("/clans/details", controllers.DetailsClan)
}
