package schemas

import (
	"encoding/json"
	"net/http"
)

type SchemaDatabaseError struct {
	Code    int
	Message string
}

type SchemaErrorResponse struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
