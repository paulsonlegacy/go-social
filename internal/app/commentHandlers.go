package app

import (
	"log"
	"net/http"
	"strconv"
	"errors"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/paulsonlegacy/go-social/internal/models"
)

// TYPES

type CreateCommentPayload struct {
	UserID    int64    `json:"user_id" validate:"required"`
	PostID    sql.NullInt64    `json:"post_id" validate:"omitempty"`
	ParentID  sql.NullInt64  `json:"parent_id" validate:"omitempty"`
	Content   string   `json:"content" validate:"required,max=10000"`
}

type UpdateCommentPayload struct {
	Content  string   `json:"content" validate:"required,max=10000"`
}


// CreateCommentHandler handles the creation of a new comment
func (app *Application) CreateCommentHandler(responseW http.ResponseWriter, request *http.Request) {

	// Define a variable to store the incoming request payload
	var payload CreateCommentPayload

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

	// Attempt to create a new comment in the database
	if newComment, err := app.Models.Comments.Create(
		ctx,
		&models.Comment{
			UserID:  payload.UserID,  // Assign UserID from payload
			PostID:  payload.PostID,  // Assign PostID from payload
			ParentID:  payload.ParentID,  // Assign ParentID from payload
			Content: payload.Content, // Assign Content from payload
		},
	); err != nil {

		// If there's an error, respond with an Internal Server Error and return
		app.StatusInternalServerError(responseW, request, err); return

	} else {

		// If successful, respond with a success message
		app.StatusCreated(responseW, request, newComment)

	}
}



// UpdatePostHandler updates a single post
func (app *Application) UpdateCommentHandler(responseW http.ResponseWriter, request *http.Request) {

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
	existingComment, err := app.Models.Comments.GetByID(ctx, id)

	if err != nil {

		app.StatusNotFound(responseW, request, errors.New("post not found")); return

	}

	log.Println("Existing comment before update: ", existingComment)

	// Define a variable to store the incoming request payload
	var payload UpdateCommentPayload

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

	if payload.Content == "" {

		if err := app.Models.Comments.Delete(ctx, id); err != nil {

			app.StatusOK(responseW, request, "comment successfully deleted"); return

		}

		app.StatusBadRequest(responseW, request, errors.New("comment cannot be empty"));  return


	} else {

		// Update comment
		existingComment.Content = payload.Content

		log.Println("Existing post after update: ", existingComment)

		// Attempt to update comment in the database
		updatedComment, err := app.Models.Comments.Update(ctx, existingComment)

		if err != nil {

			app.StatusInternalServerError(responseW, request, err); return

		}

		app.StatusUpdated(responseW, request, updatedComment)

	}

}



// DeletePostHandler deletes a single comment
func (app *Application) DeleteCommentHandler(responseW http.ResponseWriter, request *http.Request) {

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
	err = app.Models.Comments.Delete(ctx, id)

	if err != nil {

		app.StatusInternalServerError(responseW, request, err); return

	}

	app.StatusOK(responseW, request, "comment successfully deleted")

}
