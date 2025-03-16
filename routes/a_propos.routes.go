package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func AProposRoutes() {
	http.HandleFunc("/a_propos", controllers.AProposController)
}
