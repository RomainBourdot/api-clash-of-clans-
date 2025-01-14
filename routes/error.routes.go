package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func ErrorRoutes() {

	http.HandleFunc("/error", controllers.ErrorController)
}
