package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func FavoritesRoutes() {
	http.HandleFunc("/favorites", controllers.FavoriteController)
	http.HandleFunc("/favorites/list", controllers.ListFavoritesController)
	http.HandleFunc("/favorites/add", controllers.AddFavoriteController)
	http.HandleFunc("/favorites/remove", controllers.RemoveFavoriteController)
}
