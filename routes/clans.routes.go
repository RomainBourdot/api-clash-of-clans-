package routes

import "net/http"

func clansRoutes() {
	// route permettant d'accéder à la page des clans
	// la route /clans est associé au contrôleur clansController
	http.HandleFunc("/clans", controllers.ClansController)
}
