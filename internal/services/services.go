package services

import (
	"strings"
	"encoding/json"
	"net/http"
)

type JSONMessage struct {
	Message string `json:"message"` // The JSON response will have an "message" field.
}

// FUNCTIONS

// WriteJSON sends a JSON response with the given status code and data.
func WriteJSON(w http.ResponseWriter, status int, data any) error {
	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code for the response.
	w.WriteHeader(status)

	// Encode the data as JSON and send it in the response.
	return json.NewEncoder(w).Encode(data)
}

// readJSON reads JSON data from the request body and decodes it into the provided data structure.
func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error { 
	// Set a limit on the request body size (1MB) to prevent large payloads.
	maxBytes := 1_048_578  
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))  

	// Create a JSON decoder for the request body.
	decoder := json.NewDecoder(r.Body)  

	// Disallow extra fields in the JSON to prevent unexpected data.
	decoder.DisallowUnknownFields()  

	// Decode the JSON into the provided data structure and return any error.
	return decoder.Decode(data)  
}

// WriteJSONStatus sends a JSON response with a message.
func WriteJSONStatus(w http.ResponseWriter, status int, message string) error {
	// Call WriteJSON to send the error response with the given status code.
	return WriteJSON(w, status, &JSONMessage{Message: message})
}

// ConvertTags checks if the data slice contains a single string and splits it into multiple data if needed.
func ConvertStringToSlice(data []string) []string {
	if len(data) == 1 {
		return strings.Split(data[0], ",") // Convert "golang, web, api" -> ["golang", "web", "api"]
	}
	return data
}