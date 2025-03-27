package home

import "net/http"

type HomeHandler struct {}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PokéAPI Mini home page."))
}
