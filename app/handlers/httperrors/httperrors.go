package httperrors

import "net/http"

func ForbiddenHandler(w http.ResponseWriter, r *http.Request) string {
	msg := []byte("403 Forbidden")
    w.WriteHeader(http.StatusForbidden)
    w.Write(msg)

	return string(msg[:])
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) string {
	msg := []byte("404 Not Found")
    w.WriteHeader(http.StatusNotFound)
    w.Write(msg)

	return string(msg[:])
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) string {
	msg := []byte("405 Method Not Allowed")
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write(msg)

	return string(msg[:])
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) string {
	msg := []byte("500 Internal Server Error")
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(msg)

	return string(msg[:])
}
