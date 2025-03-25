package handlers

import (
	"net/http"

	"github.com/paulsonlegacy/go-social/internal/app"
	"github.com/paulsonlegacy/go-social/internal/services"
	"github.com/paulsonlegacy/go-social/internal/models"
)


type CreatePostPayload struct {
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
        services.WriteJSONError(responseW, http.StatusBadRequest, err.Error())
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
        services.WriteJSONError(responseW, http.StatusInternalServerError, err.Error())
        return
    } else {
        // If successful, respond with a success message
        services.WriteJSONSuccess(responseW, http.StatusCreated, "Post successfully created")
        return
    }
}