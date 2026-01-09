package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

var key []byte = []byte(os.Getenv("APP_JWT_KEY"))

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		log.Println("Jestem tu 1")
		if jwtToken == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(jwtToken.Value, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return key, nil
		})
		log.Println("Jestem tu 2")
		if token == nil {
			writeUnauthed(w)
			return
		}
		log.Println("Jestem tu 3")
		if err != nil && !token.Valid {
			log.Println("Parse error: " + err.Error())
			if !token.Valid {
				log.Println("Token is not valid")
			}
			writeUnauthed(w)
			return
		}
		log.Println("Jestem tu 4")
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Println(claims)
			log.Println("Jestem tutaj w auth middleware")
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	})
}
