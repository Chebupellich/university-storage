package middleware

import (
	"context"
	"fmt"
	"net/http"
	"server/models"

	"github.com/dgrijalva/jwt-go"
)

type RToken struct {
	Refresh string `json:"refresh_token"`
}

var jwtKey = []byte("tut_mogla-bit__w-a-s-h-a__REKLAMA")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		claims := &models.Claims{}

		if len(tokenStr) <= len("Bearer ") {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}
		tokenStr = tokenStr[len("Bearer "):]

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoleMiddleware(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("claims").(*models.Claims)
			if !ok || claims.Role != role {
				http.Error(w, "доступ запрещен", http.StatusForbidden)
				return
			}

			fmt.Println("Request correct")
			next.ServeHTTP(w, r)
		})
	}
}
