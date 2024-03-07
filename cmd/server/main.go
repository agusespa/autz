package main

import (
	"log"
	"net/http"

	"github.com/agusespa/autz/internal/database"
	"github.com/agusespa/autz/internal/handlers"
	"github.com/agusespa/autz/internal/repository"
	"github.com/agusespa/autz/internal/service"
)

func main() {
	db, dbErr := database.ConnectDB()
	if dbErr != nil {
		log.Fatalf("Error establishing database connection: %v", dbErr)
	}

	// TODO: get port dinamically
	port := "3001"

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewProductService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	initHandlers(authHandler)

	log.Printf("Listening on port %v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting the HTTP server: %v", err)
	}
}

func initHandlers(authHandler *handlers.AuthHandler) {
	http.HandleFunc("/authapi/register", authHandler.HandleUserRegister)
	http.HandleFunc("/authapi/login", authHandler.HandleUserLogin)
	http.HandleFunc("/authapi/user", authHandler.HandleUserDataEdit)
	http.HandleFunc("/authapi/authenticate", authHandler.HandleUserAuthentication)
	http.HandleFunc("/authapi/refresh", authHandler.HandleTokenRefresh)
	http.HandleFunc("/authapi/logout/all", authHandler.HandleUserTokensRevocation)
	http.HandleFunc("/authapi/logout", authHandler.HandleTokenRevocation)
}
