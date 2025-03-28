package app

import (
	"encoding/json"
	"log"
	"net/http"
)

// TYPES

// Response model
type HTTPResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}


// FUNCTIONS

// NewHTTPResponse initializes a response model
func (app *Application) NewHTTPResponse(status int, data any) HTTPResponse {
	switch status {

		// Handle error-related status codes
		case http.StatusBadRequest,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusForbidden,
			http.StatusUnauthorized,
			http.StatusNotFound,
			http.StatusRequestTimeout:

			// If data is an error type, return an error response with the error message
			if err, ok := data.(error); ok {
				
				// Log the error
				log.Println(err.Error())

				return HTTPResponse{
					Status:  status,
					Success: false,
					Message: err.Error(),
				}
			}

			// If data is a string, return an error response with the string as the message
			if err, ok := data.(string); ok {
				return HTTPResponse{
					Status:  status,
					Success: false,
					Message: err,
				}
			}

			// Default error response if data type is unknown
			return HTTPResponse{
				Status:  status,
				Success: false,
				Message: "Error occurred",
			}

		default:

			// If data is a string, return a success response with the message
			if successMessage, ok := data.(string); ok {
				return HTTPResponse{
					Status:  status,
					Success: true,
					Message: successMessage,
				}
			}

			// Default success response with included data
			return HTTPResponse{
				Status:  status,
				Success: true,
				Message: "Request successful",
				Data:    data,
			}
	}
}





// WriteJSON sends a JSON response with the given status code and data.
func (app *Application) WriteJSON(w http.ResponseWriter, status int, data any) error {

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code for the response.
	w.WriteHeader(status)

	// Encode the data as JSON and send it in the response.
	return json.NewEncoder(w).Encode(data)

}





// readJSON reads JSON data from the request body and decodes it into the provided data structure.
func (app *Application) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
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
