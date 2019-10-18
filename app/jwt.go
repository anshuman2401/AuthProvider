package app

import (
	"AuthProvider/models"
	util "AuthProvider/utils"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var JWTAuthentication = func(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		notAuth := map[string]bool{"/api/signup/email": true, "/api/auth/email": true}
		requestPath := req.URL.Path

		if _, ok := notAuth[requestPath]; ok {
			next.ServeHTTP(w, req)
			return
		}

		response := make(map[string]interface{})

		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" {
			response := util.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")

		if len(splitted) != 2 {
			response := util.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		jwtToken := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(jwtToken, tk, func(token *jwt.Token) (i interface{}, e error) {
			return []byte("token"), nil
		})

		if err!=nil {
			response = util.Message(false, "Malformed authentication token")
			fmt.Print(err)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		if !token.Valid {
			response = util.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		fmt.Sprintf("User %s", tk.Uidx)
		ctx := context.WithValue(req.Context(), "user", tk.UserId)

		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
