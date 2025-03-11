package pokeapi

import (
	"context"
	"encoding/json"
	"mon-go-rest/config/logging"
	"mon-go-rest/handlers/errors"
	"mon-go-rest/models"
	"mon-go-rest/mongodb"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// URIs
var (
    PokemonRe = regexp.MustCompile(`^/api/v1/pokemon/*$`)
    PokemonIdRe = regexp.MustCompile(`^/api/v1/pokemon/([0-9]{1,4})$`)
)

type PokemonHandler struct {}

func (h *PokemonHandler) ListAllPokemon(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, "pokemon", "kanto")

	if conn.Collection == nil {
		return
	}

	defer conn.CloseConnection(ctx)
	
	records, err := mongodb.GetAllRecords[models.Pokemon](ctx, *conn.Collection)

	if err != nil {
		logger.Info("HTTP Error 404: Not found.")
        errors.NotFoundHandler(w, r)
        return
    }

    jsonRecords, err := json.MarshalIndent(records, "", "  ")

    if err != nil {
        errors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecords)
}

func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, "pokemon", "kanto")
	defer conn.CloseConnection(ctx)
	
	// Extract the resource ID/slug using a regex
    matches := PokemonIdRe.FindStringSubmatch(r.URL.Path)

    // Expect matches to be length >= 2 (full string and 1 matching group)
    if len(matches) < 2 {
        errors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

	idValue, _ := strconv.Atoi(matches[1])
    pokemon, err := mongodb.GetRecordById[models.Pokemon](ctx, *conn.Collection, "_id", idValue)
    
	if err != nil {
        errors.NotFoundHandler(w, r)
		logger.Info("HTTP Error 404: Not found.")
        return
    }

    jsonRecord, err := json.MarshalIndent(pokemon, "", "  ")

    if err != nil {
        errors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecord)
}

func (h *PokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == http.MethodGet && PokemonRe.MatchString(r.URL.Path):
        h.ListAllPokemon(w, r)
        return
    case r.Method == http.MethodGet && PokemonIdRe.MatchString(r.URL.Path):
        h.GetPokemon(w, r)
        return
    default:
        return
    }
}
