package respond

import (
	"encoding/json"
	"net/http"
)

// WithJSON writes a JSON response.
func WithJSON(w http.ResponseWriter, status int, of interface{}) error {
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	err := e.Encode(of)
	if err != nil {
		http.Error(w, "error encoding error", http.StatusInternalServerError)
	}
	return err
}

// Error returned by the API.
type Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// WithError responds with a JSON error.
func WithError(w http.ResponseWriter, status int, msg string) error {
	return WithJSON(w, status, Error{
		Error:  msg,
		Status: status,
	})
}

// WithOK responds with a JSON OK.
func WithOK(w http.ResponseWriter) error {
	return WithJSON(w, http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
