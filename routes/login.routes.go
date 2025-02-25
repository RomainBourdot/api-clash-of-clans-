package routes

import (
	"groupie-tracker/login"
	"net/http"
)

func loginRoutes() {
	http.HandleFunc("/login", login.LoginController)
}
