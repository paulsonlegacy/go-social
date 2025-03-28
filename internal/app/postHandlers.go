package app

import (
	"log"
	"net/http"
	"strconv"

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

		responseData := app.NewHTTPResponse(http.StatusBadRequest, err.Error())

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

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
		responseData := app.NewHTTPResponse(http.StatusInternalServerError, err.Error())

		app.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

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

		responseData := app.NewHTTPResponse(http.StatusBadRequest, err.Error())

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	} else if posts == nil {

		responseData := app.NewHTTPResponse(http.StatusNotFound, "Resource not found")

		app.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	} else {

		responseData := app.NewHTTPResponse(http.StatusOK, posts)

		app.WriteJSON(responseW, http.StatusOK, responseData)

		return
	}
}





// FetchPostHandler fetches a single post
func (app *Application) FetchPostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := app.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	data, err := app.Models.Posts.GetByID(ctx, id)

	// Logging data
	log.Println(data)

	if data == nil {

		responseData := app.NewHTTPResponse(http.StatusNotFound, "Resource not found")

		app.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	} else if err != nil {

		responseData := app.NewHTTPResponse(http.StatusBadRequest, "Bad request")

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	} else {

		responseData := app.NewHTTPResponse(http.StatusFound, data)

		app.WriteJSON(responseW, http.StatusFound, responseData)

		return

	}

}





// UpdatePostHandler updates a single post
func (app *Application) UpdatePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract post ID from URL
	idStr := chi.URLParam(request, "id") // assuming you're using `chi` router

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := app.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Request context
	ctx := request.Context()

	// Fetch the existing post
	existingPost, err := app.Models.Posts.GetByID(ctx, id)

	if err != nil {

		responseData := app.NewHTTPResponse(http.StatusNotFound, "Post not found")

		app.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	}

	// Define a variable to store the incoming request payload
	var payload PostPayload

	// Read & parse JSON request body into payload
	if err := app.ReadJSON(responseW, request, &payload); err != nil {

		responseData := app.NewHTTPResponse(http.StatusBadRequest, "Bad request")

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

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

		responseData := app.NewHTTPResponse(http.StatusInternalServerError, err)

		app.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

	}

	responseData := app.NewHTTPResponse(http.StatusOK, "Post successfully updated")

	app.WriteJSON(responseW, http.StatusOK, responseData)

}





// DeletePostHandler deletes a single post
func (app *Application) DeletePostHandler(responseW http.ResponseWriter, request *http.Request) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := app.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		app.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	err = app.Models.Posts.Delete(ctx, id)

	if err != nil {

		responseData := app.NewHTTPResponse(http.StatusInternalServerError, "Internal server error")

		app.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

	}

	responseData := app.NewHTTPResponse(http.StatusOK, "Post successfully deleted")

	app.WriteJSON(responseW, http.StatusOK, responseData)

}
