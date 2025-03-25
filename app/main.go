package main

import (
	"errors"
	"fmt"
	"mon-go-rest/config/logging"
	"mon-go-rest/config/mongodb"
	"mon-go-rest/config/server"
	"mon-go-rest/handlers/home"
	"mon-go-rest/handlers/pokeapi"
	"net/http"
	"os"
)

func main() {
	// Logger
	logging.LoadJSONLogger()
	logger := logging.Logger
	logger.Info("Application started.")
	
	// Load environment vars
	if err := server.LoadConfig(); err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	if err := mongodb.LoadConfig(); err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	logger.Info("Environment variables loaded.")

	// Server set up
	serverUri := server.Server.Uri
	mux := http.NewServeMux()
	mux.Handle("/", &home.HomeHandler{})
    mux.Handle("/api/v1/pokemon", &pokeapi.PokemonHandler{})
    mux.Handle("/api/v1/pokemon/", &pokeapi.PokemonHandler{})

	logger.Info(fmt.Sprintf("Starting server at '%s'.", serverUri))

	logMux := logging.HttpRequestLogging(logger)(mux)
	err := http.ListenAndServe(serverUri, logMux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Error("Error: server closed.")
		os.Exit(1)
	} else if err != nil {
		logger.Error(fmt.Sprintf("Error: couldn't start server, %s\n", err))
		os.Exit(1)
	}
}
	