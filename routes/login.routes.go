package routes

import (
	"groupie-tracker/login"
	"net/http"
)

func loginRoutes() {
	http.HandleFunc("/login", login.LoginController)
	http.HandleFunc("/login/traitement", login.LoginTraitement)
}
