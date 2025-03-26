package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"mon-go-rest/config/logging"
	"mon-go-rest/handlers/httperrors"
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
			logger.Info("HTTP Error 404: Not found.")
			httperrors.NotFoundHandler(w, r)
			return
		}

		allRecords = append(allRecords, records...)
	}

	jsonRecords, err := json.MarshalIndent(allRecords, "", "  ")

    if err != nil {
        httperrors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecords)
}

func (h *PokemonHandler) ListAllPokemoByRegion(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger
	match := getRegexUrlMatch(PokemonRegionRe, r.URL.Path)

    if match == "" {
        httperrors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	records, err := getRecordsByRegion(ctx, match)
		
	if err != nil {
		logger.Info("HTTP Error 404: Not found.")
		httperrors.NotFoundHandler(w, r)
		return
	}

	jsonRecords, err := json.MarshalIndent(records, "", "  ")

    if err != nil {
        httperrors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecords)
}

func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger

	match := getRegexUrlMatch(PokemonIdRe, r.URL.Path)

    if match == "" {
        httperrors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

	idValue, err := strconv.Atoi(match)

	if err != nil {
        httperrors.NotFoundHandler(w, r)
		logger.Info("HTTP Error 404: Not found.")
        return
    }

	region := getRegionByPokemonId(idValue)

	if region == "" {
		httperrors.NotFoundHandler(w, r)
		logger.Info("HTTP Error 404: Not found.")
        return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	conn := mongodb.NewConnection(ctx, "pokemon", region)
	defer conn.CloseConnection(ctx)

    pokemon, err := mongodb.GetRecordById[models.Pokemon](ctx, *conn.Collection, "_id", idValue)
    
	if err != nil {
        httperrors.NotFoundHandler(w, r)
		logger.Info("HTTP Error 404: Not found.")
        return
    }

    jsonRecord, err := json.MarshalIndent(pokemon, "", "  ")

    if err != nil {
        httperrors.InternalServerErrorHandler(w, r)
		logger.Info("HTTP Error 500: Internal Server Error.")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRecord)
}

func (h *PokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch {
	case r.Method != http.MethodGet:
		httperrors.ForbiddenHandler(w, r)
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
		httperrors.NotFoundHandler(w, r)
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
