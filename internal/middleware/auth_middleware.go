package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go-auth/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authHeader string
		authHeader = r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No Authorization Header Provided", http.StatusUnauthorized)
			return
		}

		var tokenString string
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		var token *jwt.Token
		var err error
		token, err = auth.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing token: %v", err), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var ctx context.Context
			ctx = context.WithValue(r.Context(), "user_email", claims["email"])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid Token Claims", http.StatusUnauthorized)
		}
	})
}
