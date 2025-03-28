package home

import (
	"mon-go-rest/config/logging"
	"mon-go-rest/handlers/responses"
	"net/http"
)

type HomeHandler struct {}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := responses.OkHandler(w, r, []byte("PokéAPI Mini home page."))
	logging.Logger.Info(message)
}
