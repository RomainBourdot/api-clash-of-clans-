package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func Research() {

	http.HandleFunc("/research", controllers.ResearchData)
}
