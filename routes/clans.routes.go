package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func ClansRoutes() {
	// route permettant d'accéder à la page des clans
	// la route /clans est associé au contrôleur clansController
	http.HandleFunc("/clans", controllers.ListClans)
}
