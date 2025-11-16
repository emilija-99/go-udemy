package main

import (
	"log"
	"net/http"
	"social/internal/store"
)

// stop user to corrupt store
type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	// decode json
	log.Printf("w,r: %+v r= %+v", w, r)
	// var post store.Post
	var payload CreatePostPayload
	// json like $payload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// store contains what is in payload.
	// create actual post from payload
	post := &store.Post{
		Title:   payload.Title,
		Context: payload.Content,
		UserID:  1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
