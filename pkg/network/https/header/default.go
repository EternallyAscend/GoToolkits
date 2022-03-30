package header

import "net/http"

func SetDefaultHeaders(w http.ResponseWriter) {
}

func SetContentJsonHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}
