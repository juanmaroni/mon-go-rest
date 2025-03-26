package httperrors

import "net/http"

// HTTP Error 403
func ForbiddenHandler(w http.ResponseWriter, r *http.Request) []byte {
	msg := []byte("403 Forbidden")
    w.WriteHeader(http.StatusForbidden)
    w.Write(msg)

	return msg
}

// HTTP Error 404
func NotFoundHandler(w http.ResponseWriter, r *http.Request) []byte {
	msg := []byte("404 Not Found")
    w.WriteHeader(http.StatusNotFound)
    w.Write(msg)

	return msg
}

// HTTP Error 405
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) []byte {
	msg := []byte("405 Method Not Allowed")
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write(msg)

	return msg
}

// HTTP Error 500
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) []byte {
	msg := []byte("500 Internal Server Error")
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(msg)

	return msg
}
