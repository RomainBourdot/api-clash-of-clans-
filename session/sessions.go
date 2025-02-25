// sessions.go
package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Déclaration d'un store de sessions (ici CookieStore)
var store = sessions.NewCookieStore([]byte("votre-cle-secrete"))

// GetSession récupère la session pour la requête HTTP.
func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session-name")
}
