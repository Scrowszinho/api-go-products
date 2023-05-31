package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err = jwtauth.VerifyToken(tokenAuth, tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		decodedClaims := claims["user"].(float64)
		id := int(decodedClaims)
		ctx := context.WithValue(r.Context(), "user", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
