package app

import (
	"log"
	"net/http"
)



func (app *Application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusInternalServerError, err.Error())

	app.WriteJSON(w, http.StatusInternalServerError, responseData)

}


func (app *Application) BadRequest(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusBadRequest, err.Error())

	app.WriteJSON(w, http.StatusBadRequest, responseData)

}


func (app *Application) NotFound(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusNotFound, err.Error())

	app.WriteJSON(w, http.StatusNotFound, responseData)

}


func (app *Application) RequestTimeout(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusRequestTimeout, err.Error())

	app.WriteJSON(w, http.StatusRequestTimeout, responseData)

}


func (app *Application) Forbidden(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusForbidden, err.Error())

	app.WriteJSON(w, http.StatusForbidden, responseData)

}


func (app *Application) BadGateway(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s path: %s error: %s \n", r.Method, r.URL.Path, err.Error())

	responseData := app.NewHTTPResponse(http.StatusBadGateway, err.Error())

	app.WriteJSON(w, http.StatusBadGateway, responseData)

}