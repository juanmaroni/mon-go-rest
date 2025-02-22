package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"poke-api-mini/config"
	"poke-api-mini/handlers/home"
	"poke-api-mini/handlers/pokeapi"
)

func main() {
	// Load environment vars
	config.LoadServerConfig()
	config.LoadMongoDbConfig()

	// TODO: Setup logger
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//consoleLogger := logger.NewConsoleLogger()

	// Server set up
	mux := http.NewServeMux()
	mux.Handle("/", &home.HomeHandler{})
    mux.Handle("/api/v1/pokemon", &pokeapi.PokemonHandler{})
    mux.Handle("/api/v1/pokemon/", &pokeapi.PokemonHandler{})

	jsonLogger.Info("Server started.")
	defer jsonLogger.Info("Server shut down.")

	err := http.ListenAndServe(config.Server.Uri, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Error: server closed.\n") // Log
	} else if err != nil {
		fmt.Printf("Error: couldn't start server, %s\n", err) // Log
		os.Exit(1)
	}
}
	