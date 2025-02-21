package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"poke-api-mini/handlers/errors"
	"poke-api-mini/models"
	"poke-api-mini/mongodb"
	"regexp"
	"strconv"
	"time"
)

// TODO: Environment vars
const MongoDbUri string = "mongodb://localhost:27017"

// URIs
var (
    PokemonRe = regexp.MustCompile(`^/api/v1/pokemon/*$`)
    PokemonIdRe = regexp.MustCompile(`^/api/v1/pokemon/([0-9]{1,4})$`)
)

type PokemonHandler struct {}

func (h *PokemonHandler) ListAllPokemon(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, MongoDbUri, "pokemon", "kanto")
	fmt.Println(conn.Connected) // Log
	defer conn.CloseConnection(ctx)
	
	records, err := mongodb.GetAllRecords[models.Pokemon](ctx, *conn.Collection)

	if err != nil {
        errors.NotFoundHandler(w, r)
        return
    }

    jsonRecords, err := json.MarshalIndent(records, "", "  ")

    if err != nil {
        errors.InternalServerErrorHandler(w, r)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecords)
}

func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, MongoDbUri, "pokemon", "kanto")
	fmt.Println(conn.Connected) // Log
	defer conn.CloseConnection(ctx)
	
	// Extract the resource ID/slug using a regex
    matches := PokemonIdRe.FindStringSubmatch(r.URL.Path)

    // Expect matches to be length >= 2 (full string and 1 matching group)
    if len(matches) < 2 {
        errors.InternalServerErrorHandler(w, r)
        return
    }

	idValue, _ := strconv.Atoi(matches[1])
    pokemon, err := mongodb.GetRecordById[models.Pokemon](ctx, *conn.Collection, "_id", idValue)
    
	if err != nil {
        errors.NotFoundHandler(w, r)
        return
    }

    jsonRecord, err := json.MarshalIndent(pokemon, "", "  ")

    if err != nil {
        errors.InternalServerErrorHandler(w, r)
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
