package middlewares

import (
	"log"
	"context"
	"net/http"
)

func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("user_jwt")
		if err != nil {
			log.Printf("Error getting cookie. Probably not logged in.")
			next.ServeHTTP(w, r)
		}

		token := cookie.Value
		ctx := context.WithValue(r.Context(), "userNanoId", token)
		r.Context()

		next.ServeHTTP(w, ctx)
	})
}
