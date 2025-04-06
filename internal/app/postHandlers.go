package app

import (
	"log"
	"net/http"
	"strconv"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulsonlegacy/go-social/internal/models"
)


// TYPES
type CreatePostPayload struct {
	UserID    int64    `json:"user_id" validate:"required"`
	Title     string   `json:"title" validate:"required,min=10"`
	Content   string   `json:"content" validate:"required,min=10,max=10000"`
	Tags      []string `json:"tags" validate:"omitempty"`
}

type UpdatePostPayload struct {
	UserID    int64    `json:"user_id" validate:"omitempty"`
	Title     string   `json:"title" validate:"omitempty,min=10"`
	Content   string   `json:"content" validate:"omitempty,min=10,max=10000"`
	Tags      []string `json:"tags" validate:"omitempty"`
}


// FUNCTIONS

// Index handler
func (app *Application) HomeHandler(responseW http.ResponseWriter, request *http.Request) {
	responseW.Write([]byte("welcome"))
}





// CreatePostHandler handles the creation of a new post
func (app *Application) CreatePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Define a variable to store the incoming request payload
	var payload CreatePostPayload

	// Read and parse the JSON request body into the payload struct
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Validating payload
	if err := Validate.Struct(payload); err != nil {
		
		app.StatusBadRequest(responseW, request, err); return

	}

	// Get the request's context
	ctx := request.Context()

	// Attempt to create a new post in the database
	if newPost, err := app.Models.Posts.Create(
		ctx,
		&models.Post{
			UserID:  payload.UserID,  // Assign UserID from payload
			Title:   payload.Title,   // Assign Title from payload
			Content: payload.Content, // Assign Content from payload
			Tags:    payload.Tags,    // Assign Tags from payload
		},
	); err != nil {

		// If there's an error, respond with an Internal Server Error and return
		app.StatusInternalServerError(responseW, request, err); return

	} else {

		// If successful, respond with a success message
		app.StatusCreated(responseW, request, newPost)

	}
}





// FetchPostsHandler fetches all posts without pagination
func (app *Application) FetchPostsHandler(responseW http.ResponseWriter, request *http.Request) {

	// Get the request's context
	ctx := request.Context()

	posts, err := app.Models.Posts.GetAll(ctx)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err); return


	} else if posts == nil {

		app.StatusNotFound(responseW, request, errors.New("resource not found")); return

	} else {

		app.StatusOK(responseW, request, posts); return

	}
}





// FetchPostHandler fetches a single post
func (app *Application) FetchPostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	data, err := app.Models.Posts.GetByID(ctx, id)

	// Logging data
	log.Println(data)

	if data == nil {

		app.StatusNotFound(responseW, request, errors.New("post not found")); return

	} else if err != nil {

		app.StatusInternalServerError(responseW, request, err); return

	} else {

		app.StatusFound(responseW, request, data); return

	}

}





// UpdatePostHandler updates a single post
func (app *Application) UpdatePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract post ID from URL
	idStr := chi.URLParam(request, "id") // assuming you're using `chi` router

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Request context
	ctx := request.Context()

	// Fetch the existing post
	existingPost, err := app.Models.Posts.GetByID(ctx, id)

	if err != nil {

		app.StatusNotFound(responseW, request, errors.New("post not found")); return

	}

	log.Println("Existing post before update: ", existingPost)

	// Define a variable to store the incoming request payload
	var payload UpdatePostPayload

	// Read & parse JSON request body into payload
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusInternalServerError(responseW, request, err); return

	}

	// Validating payload
	if err := Validate.Struct(payload); err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	log.Println("Payload: ", payload)

	// Update only fields that are provided in the request

	if payload.UserID != 0 {
		existingPost.UserID = payload.UserID
	}

	if payload.Title != "" {
		existingPost.Title = payload.Title
	}

	if payload.Content != "" {
		existingPost.Content = payload.Content
	}

	if payload.Tags != nil { // Check if Tags are provided
		existingPost.Tags = payload.Tags
	}

	log.Println("Existing post after update: ", existingPost)

	// Attempt to update post in the database
	updatedPost, err := app.Models.Posts.Update(ctx, existingPost)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err); return

	}

	app.StatusUpdated(responseW, request, updatedPost)

}





// DeletePostHandler deletes a single post
func (app *Application) DeletePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err); return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	err = app.Models.Posts.Delete(ctx, id)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err); return

	}

	app.StatusOK(responseW, request, "Post successfully deleted")

}
