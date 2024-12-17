package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func AccueilRoutes() {
	// route permettant d'accéder à une page d'accueil
	// la route /accueil est associé au contrôleur accueilController
	http.HandleFunc("/accueil", controllers.AccueilController)
}
