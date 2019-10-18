package main

import (
	"AuthProvider/app"
	"AuthProvider/controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JWTAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/api/signup/email", controller.CreateAccount).Methods("POST")
	router.HandleFunc("/api/auth/login", controller.AuthenticateEmail).Methods("POST")
	router.HandleFunc("/api/auth/signout", controller.Signout).Methods("POST")

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print("Error in connecting to server", err)
	}
}
