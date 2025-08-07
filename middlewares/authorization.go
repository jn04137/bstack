package middlewares

import (
	"log"
	"context"
	"net/http"

	"com/bstack/services"
)

func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("user_jwt")
		if err != nil {
			log.Printf("Error getting cookie. Probably not logged in.")
			next.ServeHTTP(w, r)
			return
		}

		token := cookie.Value
		username, nanoId, err := services.ValidateToken(token)
		if err != nil {
			log.Printf("Error getting username and nanoid from user: %v", err.Error())
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		ctx = context.WithValue(r.Context(), "userNanoId", nanoId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
