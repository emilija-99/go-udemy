package main

import (
	"errors"
	"log"
	"net/http"
	"social/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// stop user to corrupt store
// error occurs if you sed validate:"required, max=100" -> if there are some spaces
type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	// decode json
	log.Printf("w,r: %+v r= %+v", w, r)
	// var post store.Post
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.statusBadRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// store contains what is in payload.
	// create actual post from payload
	post := &store.Post{
		Title:   payload.Title,
		Context: payload.Content,
		UserID:  1,
	}

	err := errors.New("Payload is empty! (:")
	if payload.Content == "" {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postId")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}
	ctx := r.Context()
	post, err := app.store.Posts.GetById(ctx, id)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.statusNotFound(w, r, err)
		default:
			app.statusInternlServerError(w, r, err)
		}
		return
	}

	if err == nil {
		writeJSON(w, http.StatusOK, post)
	}
	log.Printf("Post: %+v", post)
}
