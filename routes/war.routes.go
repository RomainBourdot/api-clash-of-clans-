package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func WarRoutes() {
	http.HandleFunc("/wars", controllers.ListWars)
	http.HandleFunc("/wars/details", controllers.WarDetails)
}
