package routes

import (
	"fmt"
	"net/http"
)

func InitServe() {

	AccueilRoutes()
	ErrorRoutes()
	ClansRoutes()
	Research()
	FavoritesRoutes()
	loginRoutes()
	RegisterRoutes()
	WarRoutes()

	fmt.Println("Le serveur est opérationel : http://localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}
