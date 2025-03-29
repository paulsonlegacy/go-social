package app

import (
	"log"
	"net/http"
)


// ERROR METHODS


// Internal server error
func (app *Application) StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusInternalServerError, "internal server error")

	app.WriteJSON(w, http.StatusInternalServerError, responseData)

}

// Bad request
func (app *Application) StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusBadRequest, "bad request")

	app.WriteJSON(w, http.StatusBadRequest, responseData)

}

// Not found
func (app *Application) StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusNotFound, "not found")

	app.WriteJSON(w, http.StatusNotFound, responseData)

}

// Request timeout
func (app *Application) StatusRequestTimeout(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusRequestTimeout, "request timeout")

	app.WriteJSON(w, http.StatusRequestTimeout, responseData)

}

// Forbidden
func (app *Application) StatusForbidden(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusForbidden, "forbidden")

	app.WriteJSON(w, http.StatusForbidden, responseData)

}

// Bad gateway
func (app *Application) StatusBadGateway(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusBadGateway, "bad gateway")

	app.WriteJSON(w, http.StatusBadGateway, responseData)

}


// SUCCESS METHODS

// Status OK
func (app *Application) StatusOK(w http.ResponseWriter, r *http.Request, data any) {

	responseData := app.NewHTTPResponse(http.StatusOK, data)

	app.WriteJSON(w, http.StatusOK, responseData)

}

// Status created
func (app *Application) StatusCreated(w http.ResponseWriter, r *http.Request, data any) {

	responseData := app.NewHTTPResponse(http.StatusCreated, data)

	app.WriteJSON(w, http.StatusCreated, responseData)

}

// Status found
func (app *Application) StatusFound(w http.ResponseWriter, r *http.Request, data any) {

	responseData := app.NewHTTPResponse(http.StatusFound, data)

	app.WriteJSON(w, http.StatusFound, responseData)

}

// Status updated
func (app *Application) StatusUpdated(w http.ResponseWriter, r *http.Request, data any) {

	responseData := app.NewHTTPResponse(http.StatusOK, data)

	app.WriteJSON(w, http.StatusOK, responseData)

}