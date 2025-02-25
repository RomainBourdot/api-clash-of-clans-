package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/register", controllers.RegisterController)
}
