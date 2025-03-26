package handlers

import (
	"log"
	"strconv"
	"net/http"
	"github.com/paulsonlegacy/go-social/internal/app"
	"github.com/paulsonlegacy/go-social/internal/services"
	"github.com/paulsonlegacy/go-social/internal/models"
	"github.com/go-chi/chi/v5"
)


type CreatePostPayload struct {
	ID  int64   `json:"id"`
	UserID  int64   `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}


func HomeHandler(responseW http.ResponseWriter, request *http.Request) {
	responseW.Write([]byte("welcome"))
}

// CreatePostHandler handles the creation of a new post
func CreatePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {
    // Define a variable to store the incoming request payload
    var payload CreatePostPayload

    // Read and parse the JSON request body into the payload struct
    if err := services.ReadJSON(responseW, request, &payload); err != nil {
        // If there's an error, respond with a Bad Request error and return
        services.WriteJSONStatus(responseW, http.StatusBadRequest, err.Error())
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
        services.WriteJSONStatus(responseW, http.StatusInternalServerError, err.Error())
        return
    } else {
        // If successful, respond with a success message
        services.WriteJSONStatus(responseW, http.StatusCreated, "Post successfully created")
        return
    }
}

func FetchPostsHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {
	// Get the request's context
	ctx := request.Context()
	
	posts, err := app.Models.Posts.GetAll(ctx)

	if err != nil {
		services.WriteJSONStatus(responseW, http.StatusBadRequest, err.Error())
		return
	} else if posts == nil {
		services.WriteJSONStatus(responseW, http.StatusNotFound, "Resource not found")
		return
	} else {
		services.WriteJSON(responseW, http.StatusOK, posts)
		return
	}
}

func FetchPostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {
	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		services.WriteJSONStatus(responseW, http.StatusBadRequest, "Invalid post ID")
		return
	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	data, err := app.Models.Posts.GetByID(ctx, id);

	// Logging data
	log.Println(data)

    if data == nil {
		services.WriteJSONStatus(responseW, http.StatusNotFound, "Resource not found")
		return
	} else if err != nil {
		services.WriteJSONStatus(responseW, http.StatusBadRequest, "Bad request")
		return
	} else {
		services.WriteJSON(responseW, http.StatusFound, data)
		return
	}
	
}

func UpdatePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {
	// Define a variable to store the incoming request payload
	var payload CreatePostPayload

	// Read and parse the JSON request body into the payload struct
	if err := services.ReadJSON(responseW, request, &payload); err != nil {
		// If there's an error, respond with a Bad Request error and return
		services.WriteJSONStatus(responseW, http.StatusBadRequest, err.Error())
		return
	}

	// Get the request's context
	ctx := request.Context()

    // Attempt to update post in the database
    if err := app.Models.Posts.Update(
        ctx,
        &models.Post{
            Title:   payload.Title,   // Assign Title from payload
            Content: payload.Content, // Assign Content from payload
            Tags:    payload.Tags,    // Assign Tags from payload
			ID: payload.ID, 			  // Assign ID from payload
        },
    ); err != nil {
        // If there's an error, respond with an Internal Server Error and return
        services.WriteJSONStatus(responseW, http.StatusInternalServerError, err.Error())
        return
    } else {
        // If successful, respond with a success message
        services.WriteJSONStatus(responseW, http.StatusCreated, "Post successfully updated")
        return
    }
}

func DeletePostHandler(responseW http.ResponseWriter, request *http.Request, app *app.Application) {
	// Extract the id as a string
	idStr := chi.URLParam(request, "id")

	// Convert id to int64
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		services.WriteJSONStatus(responseW, http.StatusBadRequest, "Invalid post ID")
		return
	}

	// Get the request's context
	ctx := request.Context()

	// Query DB through post model method
	err = app.Models.Posts.Delete(ctx, id);

	if err != nil {
		services.WriteJSONStatus(responseW, http.StatusBadRequest, "Bad request")
		return
	}

	services.WriteJSONStatus(responseW, http.StatusOK, "Post successfully deleted")
}