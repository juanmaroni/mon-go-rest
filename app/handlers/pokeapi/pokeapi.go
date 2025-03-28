package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"mon-go-rest/config/logging"
	"mon-go-rest/handlers/responses"
	"mon-go-rest/models"
	"mon-go-rest/mongodb"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// URIs
var (
    PokemonRe = regexp.MustCompile(`^/api/v1/pokemon/*$`) // Get all
	PokemonRegionRe = regexp.MustCompile(`^/api/v1/pokemon/(kanto|johto)$`) // Get all by region
    PokemonIdRe = regexp.MustCompile(`^/api/v1/pokemon/([0-9]{1,4})$`) // Get one by Pokédex number (id)

	Regions = []string{"kanto", "johto"}
)

type PokemonHandler struct {}

func (h *PokemonHandler) ListAllPokemon(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	allRecords := []models.Pokemon{}

	for _, region := range Regions {
		records, err := getRecordsByRegion(ctx, region)
		
		if err != nil {
			message := responses.NotFoundHandler(w, r)
			logger.Info(message)
			return
		}

		allRecords = append(allRecords, records...)
	}

	jsonRecords, err := json.MarshalIndent(allRecords, "", "  ")

    if err != nil {
        message := responses.InternalServerErrorHandler(w, r)
		logger.Error(message)
        return
    }

	message:= responses.OkHandler(w, r, jsonRecords)
	logger.Info(message)
}

func (h *PokemonHandler) ListAllPokemoByRegion(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger
	match := getRegexUrlMatch(PokemonRegionRe, r.URL.Path)

    if match == "" {
        message := responses.InternalServerErrorHandler(w, r)
		logger.Error(message)
        return
    }

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	records, err := getRecordsByRegion(ctx, match)
		
	if err != nil {
		message := responses.NotFoundHandler(w, r)
		logger.Info(message)
		return
	}

	jsonRecords, err := json.MarshalIndent(records, "", "  ")

    if err != nil {
        message := responses.InternalServerErrorHandler(w, r)
		logger.Error(message)
        return
    }

    message := responses.OkHandler(w, r, jsonRecords)
	logger.Info(message)
}

func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger

	match := getRegexUrlMatch(PokemonIdRe, r.URL.Path)

    if match == "" {
        message := responses.InternalServerErrorHandler(w, r)
		logger.Error(message)
        return
    }

	idValue, err := strconv.Atoi(match)

	if err != nil {
        message := responses.NotFoundHandler(w, r)
		logger.Info(message)
        return
    }

	region := getRegionByPokemonId(idValue)

	if region == "" {
		message := responses.NotFoundHandler(w, r)
		logger.Info(message)
        return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, "pokemon", region)
	defer conn.CloseConnection(ctx)

    pokemon, err := mongodb.GetRecordById[models.Pokemon](ctx, *conn.Collection, "_id", idValue)
    
	if err != nil {
        message := responses.NotFoundHandler(w, r)
		logger.Info(message)
        return
    }

    jsonRecord, err := json.MarshalIndent(pokemon, "", "  ")

    if err != nil {
        message := responses.InternalServerErrorHandler(w, r)
		logger.Error(message)
        return
    }

    message := responses.OkHandler(w, r, jsonRecord)
	logger.Info(message)
}

func (h *PokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodGet:
		message := responses.MethodNotAllowedHandler(w, r)
		logging.Logger.Info(message)
        return
    case r.Method == http.MethodGet && PokemonRe.MatchString(r.URL.Path):
        h.ListAllPokemon(w, r)
        return
	case r.Method == http.MethodGet && PokemonRegionRe.MatchString(r.URL.Path):
        h.ListAllPokemoByRegion(w, r)
        return
    case r.Method == http.MethodGet && PokemonIdRe.MatchString(r.URL.Path):
        h.GetPokemon(w, r)
        return
    default:
		message := responses.NotFoundHandler(w, r)
		logging.Logger.Info(message)
        return
    }
}

func getRecordsByRegion(ctx context.Context, region string) ([]models.Pokemon, error) {
	conn := mongodb.NewConnection(ctx, "pokemon", region)
	defer conn.CloseConnection(ctx)

	if conn.Collection == nil {
		// Collection is missing, cannot give a proper response
		panic(fmt.Sprintf("Error: couldn't find collection '%s'.", region))
	}

	return mongodb.GetAllRecords[models.Pokemon](ctx, *conn.Collection)
}

func getRegionByPokemonId(id int) string {
	if id >= 1 && id <= 151 {
		return "kanto"
	} else if id >= 152 && id <= 251 {
		return "johto"
	} else {
		return ""
	}
}

func getRegexUrlMatch(re *regexp.Regexp, url string) string {
	// Extract the resource ID/slug using a regex
    matches := re.FindStringSubmatch(url)

    // Expect matches to be length == 2 (full string and 1 matching group)
    if len(matches) < 2 {
        return ""
    }

	return matches[1]
}
