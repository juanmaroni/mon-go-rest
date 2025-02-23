package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"poke-api-mini/config"
	"poke-api-mini/handlers/home"
	"poke-api-mini/handlers/pokeapi"
)

func main() {
	// Logger
	config.LoadJSONLogger()
	logger := config.Logger
	logger.Info("Application started.")
	
	// Load environment vars
	config.LoadServerConfig()
	config.LoadMongoDbConfig()
	logger.Info("Environment variables loaded.")

	// Server set up
	serverUri := config.Server.Uri
	mux := http.NewServeMux()
	mux.Handle("/", &home.HomeHandler{})
    mux.Handle("/api/v1/pokemon", &pokeapi.PokemonHandler{})
    mux.Handle("/api/v1/pokemon/", &pokeapi.PokemonHandler{})

	logger.Info(fmt.Sprintf("Starting server at '%s'.", serverUri))
	defer logger.Info("Server shut down.")

	err := http.ListenAndServe(serverUri, mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Error("Error: server closed.")
		os.Exit(1)
	} else if err != nil {
		logger.Error(fmt.Sprintf("Error: couldn't start server, %s\n", err))
		os.Exit(1)
	}
}
	