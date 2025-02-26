package routes

import (
	register "groupie-tracker/enregistrement"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/register", register.RegisterController)
	http.HandleFunc("/register/traitement", register.RegisterTraitement)
}
