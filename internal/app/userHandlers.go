package app

import (
	"log"
	"strconv"
	"net/http"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulsonlegacy/go-social/internal/models"
)


// TYPES
type CreateUserPayload struct {
	FirstName string  `json:"first_name" validate:"required,max=30"`
	LastName string  `json:"last_name" validate:"required,max=30"`
	Email string      `json:"email" validate:"required,email,max=100"`
	Username string   `json:"username" validate:"required,max=30"`
	Password string   `json:"password" validate:"required,min=8"`
}

type UpdateUserPayload struct {
	FirstName string  `json:"first_name" validate:"omitempty,max=30"`
	LastName string  `json:"last_name" validate:"omitempty,max=30"`
	Email string      `json:"email" validate:"omitempty,email,max=100"`
	Username string   `json:"username" validate:"omitempty,max=30"`
	Password string   `json:"password" validate:"omitempty,min=8"`
}


// FUNCTIONS

// CreateUserHandler handles the creation of a new user
func (app *Application) CreateUserHandler(responseW http.ResponseWriter, request *http.Request) {

	// Define a variable to store the incoming request payload
	var payload CreateUserPayload

	// Read and parse the JSON request body into the payload struct
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Validate payload
	if err := Validate.Struct(payload); err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Get the request's context
	ctx := request.Context()

	// Attempt to create a new user in the database
	if err := app.Models.Users.Create(
		ctx,
		&models.User{
			FirstName:  payload.FirstName,
			LastName:   payload.LastName,
			Username: payload.Username,
			Email:    payload.Email,
			Password: payload.Password,
		},
	); err != nil {

		// If there's an error, respond with an Internal Server Error and return
		app.StatusInternalServerError(responseW, request, err); return

	} else {

		// If successful, respond with a success message
		app.StatusCreated(responseW, request, "User successfully created")

	}
}


// FetchUsersHandler fetches all users without pagination
func (app *Application) FetchUsersHandler(responseW http.ResponseWriter, request *http.Request) {

	// Get the request's context
	ctx := request.Context()

	users, err := app.Models.Users.GetAll(ctx)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err); return


	} else if users == nil {

		app.StatusNotFound(responseW, request, errors.New("resource not found")); return

	} else {

		app.StatusOK(responseW, request, users)

	}
}



// FetchUserHandler fetches a single user
func (app *Application) FetchUserHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through user model method
	data, err := app.Models.Users.GetByID(ctx, id)

	// Logging data
	log.Println(data)

	if data == nil {

		app.StatusNotFound(responseW, request, errors.New("user not found"))

	} else if err != nil {

		app.StatusInternalServerError(responseW, request, err)

	} else {

		app.StatusFound(responseW, request, data)

	}

}





// UpdateuserHandler updates a single user
func (app *Application) UpdateUserHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract user ID from URL
	idStr := chi.URLParam(request, "id") // assuming you're using `chi` router

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Request context
	ctx := request.Context()

	// Fetch the existing user
	existingUser, err := app.Models.Users.GetByID(ctx, id)

	if err != nil {

		app.StatusNotFound(responseW, request, errors.New("user not found"))

	}

	log.Println("Existing user: ", existingUser)

	// Define a variable to store the incoming request payload
	var payload UpdateUserPayload

	// Read & parse JSON request body into payload
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusInternalServerError(responseW, request, err)

	}

	log.Println("Payload: ", payload)

	// Validate payload
	if err := Validate.Struct(payload); err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Update only fields that are provided in the request

	if payload.FirstName != "" {
		existingUser.FirstName = payload.FirstName
	}

	if payload.LastName != "" {
		existingUser.LastName = payload.LastName
	}

	if payload.Username != "" {
		existingUser.Username = payload.Username
	}

	if payload.Email != "" {
		existingUser.Email = payload.Email
	}

	if payload.Password != "" {
		existingUser.Password = payload.Password
	}

	// Attempt to update user in the database
	if err := app.Models.Users.Update(ctx, existingUser); err != nil {

		app.StatusInternalServerError(responseW, request, err)

	}

	app.StatusUpdated(responseW, request, "user successfully updated")

}





// DeleteUserHandler deletes a single user
func (app *Application) DeleteUserHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through user model method
	err = app.Models.Users.Delete(ctx, id)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err)

	}

	app.StatusOK(responseW, request, "user successfully deleted")

}