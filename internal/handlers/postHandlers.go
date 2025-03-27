package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/paulsonlegacy/go-social/internal/app"
	"github.com/paulsonlegacy/go-social/internal/models"
	"github.com/paulsonlegacy/go-social/internal/services"
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
func HomeHandler(responseW http.ResponseWriter, request *http.Request) {
	responseW.Write([]byte("welcome"))
}





// CreatePostHandler handles the creation of a new post
func CreatePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {

	// Define a variable to store the incoming request payload
	var payload PostPayload

	// Read and parse the JSON request body into the payload struct
	if err := services.ReadJSON(responseW, request, &payload); err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, err.Error())

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

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
		responseData := services.NewHTTPResponse(http.StatusInternalServerError, err.Error())

		services.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

	} else {

		// If successful, respond with a success message
		responseData := services.NewHTTPResponse(http.StatusCreated, "Post successfully created")

		services.WriteJSON(responseW, http.StatusCreated, responseData)

		return
	}
}





// FetchPostsHandler fetches all posts without pagination
func FetchPostsHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {

	// Get the request's context
	ctx := request.Context()

	posts, err := app.Models.Posts.GetAll(ctx)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, err.Error())

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	} else if posts == nil {

		responseData := services.NewHTTPResponse(http.StatusNotFound, "Resource not found")

		services.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	} else {

		responseData := services.NewHTTPResponse(http.StatusOK, posts)

		services.WriteJSON(responseW, http.StatusOK, responseData)

		return
	}
}





// FetchPostHandler fetches a single post
func FetchPostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	data, err := app.Models.Posts.GetByID(ctx, id)

	// Logging data
	log.Println(data)

	if data == nil {

		responseData := services.NewHTTPResponse(http.StatusNotFound, "Resource not found")

		services.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	} else if err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, "Bad request")

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	} else {

		responseData := services.NewHTTPResponse(http.StatusFound, data)

		services.WriteJSON(responseW, http.StatusFound, responseData)

		return

	}

}





// UpdatePostHandler updates a single post
func UpdatePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {

	// Extract post ID from URL
	idStr := chi.URLParam(request, "id") // assuming you're using `chi` router

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Request context
	ctx := request.Context()

	// Fetch the existing post
	existingPost, err := app.Models.Posts.GetByID(ctx, id)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusNotFound, "Post not found")

		services.WriteJSON(responseW, http.StatusNotFound, responseData)

		return

	}

	// Define a variable to store the incoming request payload
	var payload PostPayload

	// Read & parse JSON request body into payload
	if err := services.ReadJSON(responseW, request, &payload); err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, "Bad request")

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

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

		responseData := services.NewHTTPResponse(http.StatusInternalServerError, err)

		services.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

	}

	responseData := services.NewHTTPResponse(http.StatusOK, "Post successfully updated")

	services.WriteJSON(responseW, http.StatusOK, responseData)

}





// DeletePostHandler deletes a single post
func DeletePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {

	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusBadRequest, "Invalid post ID")

		services.WriteJSON(responseW, http.StatusBadRequest, responseData)

		return

	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	err = app.Models.Posts.Delete(ctx, id)

	if err != nil {

		responseData := services.NewHTTPResponse(http.StatusInternalServerError, "Internal server error")

		services.WriteJSON(responseW, http.StatusInternalServerError, responseData)

		return

	}

	responseData := services.NewHTTPResponse(http.StatusOK, "Post successfully deleted")

	services.WriteJSON(responseW, http.StatusOK, responseData)

}
