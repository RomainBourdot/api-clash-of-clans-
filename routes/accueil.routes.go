package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func AccueilRoutes() {

	http.HandleFunc("/", controllers.AccueilController)
}
