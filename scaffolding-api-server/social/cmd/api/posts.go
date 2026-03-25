package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"social/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type postKey string

const postCtx postKey = "post"

// stop user to corrupt store
// error occurs if you sed validate:"required, max=100" -> if there are some spaces
type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Context string   `json:"context" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}
type UpdatePost struct {
	Context *string   `json:"context"`
	Title   *string   `json:"title"`
	Tags    *[]string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	// decode json
	log.Printf("w,r: %+v r= %+v", w, r)
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
		Context: payload.Context,
		Tags:    payload.Tags,
		UserID:  1,
	}

	err := errors.New("Payload is empty! (:")
	if payload.Context == "" {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)
	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)
	if err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}

	post.Comments = comments
	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.statusInternlServerError(w, r, err)
	}

	ctx := r.Context()
	app.store.Posts.Delete(ctx, id)

	if err != nil {
		app.statusBadRequest(w, r, err)
	}
	app.jsonResponse(w, http.StatusOK, "OK")
}

func (app *application) patchPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload store.UpdatePost
	if err := readJSON(w, r, &payload); err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Context != nil {
		post.Context = *payload.Context
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	ctx := r.Context()

	if err := app.store.Posts.Patch(ctx, post); err != nil {
		app.statusInternlServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.statusInternlServerError(w, r, err)
	}

}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
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

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	log.Printf("r.context: %s", r.Context())
	post, _ := r.Context().Value(postCtx).(*store.Post)
	log.Printf("post,: %+v", post)
	return post
}
