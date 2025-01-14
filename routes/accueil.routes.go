package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func AccueilRoutes() {

	http.HandleFunc("/accueil", controllers.AccueilController)
}
