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
type PostPayload struct {
	UserID    int64    `json:"user_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}


// FUNCTIONS

// Index handler
func (app *Application) HomeHandler(responseW http.ResponseWriter, request *http.Request) {
	responseW.Write([]byte("welcome"))
}





// CreatePostHandler handles the creation of a new post
func (app *Application) CreatePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Define a variable to store the incoming request payload
	var payload PostPayload

	// Read and parse the JSON request body into the payload struct
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Get the request's context
	ctx := request.Context()

	// Attempt to create a new post in the database
	if err := app.Models.Posts.Create(
		ctx,
		&models.Post{
			UserID:  payload.UserID,  // Assign UserID from payload
			Title:   payload.Title,   // Assign Title from payload
			Content: payload.Content, // Assign Content from payload
			Tags:    payload.Tags,    // Assign Tags from payload
		},
	); err != nil {

		// If there's an error, respond with an Internal Server Error and return
		app.StatusInternalServerError(responseW, request, err)

	} else {

		// If successful, respond with a success message
		responseData := app.NewHTTPResponse(http.StatusCreated, "Post successfully created")

		app.WriteJSON(responseW, http.StatusCreated, responseData)

		return
	}
}





// FetchPostsHandler fetches all posts without pagination
func (app *Application) FetchPostsHandler(responseW http.ResponseWriter, request *http.Request) {

	// Get the request's context
	ctx := request.Context()

	posts, err := app.Models.Posts.GetAll(ctx)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err)


	} else if posts == nil {

		app.StatusNotFound(responseW, request, errors.New("resource not found"))

	} else {

		app.StatusOK(responseW, request, posts)

	}
}





// FetchPostHandler fetches a single post
func (app *Application) FetchPostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	data, err := app.Models.Posts.GetByID(ctx, id)

	// Logging data
	log.Println(data)

	if data == nil {

		app.StatusNotFound(responseW, request, errors.New("post not found"))

	} else if err != nil {

		app.StatusInternalServerError(responseW, request, err)

	} else {

		app.StatusFound(responseW, request, data)

	}

}





// UpdatePostHandler updates a single post
func (app *Application) UpdatePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract post ID from URL
	idStr := chi.URLParam(request, "id") // assuming you're using `chi` router

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Request context
	ctx := request.Context()

	// Fetch the existing post
	existingPost, err := app.Models.Posts.GetByID(ctx, id)

	if err != nil {

		app.StatusNotFound(responseW, request, errors.New("post not found"))

	}

	log.Println("Existing post: ", existingPost)

	// Define a variable to store the incoming request payload
	var payload PostPayload

	// Read & parse JSON request body into payload
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		app.StatusInternalServerError(responseW, request, err)

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

	// Attempt to update post in the database
	if err := app.Models.Posts.Update(ctx, existingPost); err != nil {

		app.StatusInternalServerError(responseW, request, err)

	}

	app.StatusUpdated(responseW, request, "Post successfully updated")

}





// DeletePostHandler deletes a single post
func (app *Application) DeletePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		app.StatusBadRequest(responseW, request, err)

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	err = app.Models.Posts.Delete(ctx, id)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err)

	}

	app.StatusOK(responseW, request, "Post successfully deleted")

}
