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

// stop user to corrupt store
// error occurs if you sed validate:"required, max=100" -> if there are some spaces
type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}
type UpdatePost struct {
	Context *string   `json:"content"`
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
		Context: payload.Content,
		Tags:    payload.Tags,
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
	log.Printf("parse:%s", chi.URLParam(r, "postId"))

	id, err := strconv.ParseInt(idParam, 10, 64)
	log.Printf("parse:%d", chi.URLParam(r, "postId"))

	if err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}
	ctx := r.Context()
	post, err := app.store.Posts.GetById(ctx, id)
	if err != nil {
		log.Printf("GET BY ID: %x", id)
		app.statusInternlServerError(w, r, err)
	}

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.statusNotFound(w, r, err)
		default:
			app.statusInternlServerError(w, r, err)
		}
		return
	}

	comments, err := app.store.Comments.GetByPostID(ctx, id)

	if err != nil {
		app.statusInternlServerError(w, r, err)
		return
	}

	if err == nil {
		post.Comments = comments
		log.Printf("Post: %+v", post)
		writeJSON(w, http.StatusOK, post)
	}

}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postId")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.statusInternlServerError(w, r, err)
	}

	ctx := r.Context()
	app.store.Posts.Delete(ctx, id)

	if err != nil {
		app.statusBadRequest(w, r, err)
	}
	writeJSON(w, http.StatusOK, "OK")
}

// func (app *application) patchPostHandler(w http.ResponseWriter, r *http.Request) {
// 	idParam := chi.URLParam(r, "postId")
// 	id, err := strconv.ParseInt(idParam, 10, 64)

// 	if err != nil {
// 		app.statusInternlServerError(w, r, err)
// 	}

// 	var input store.UpdatePost
// 	err = json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	err = app.store.Posts.Patch(r.Context(), id, &input)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, store.ErrNotFound):
// 			app.statusNotFound(w, r, err)
// 		default:
// 			app.statusInternlServerError(w, r, err)
// 		}
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

func (app *application) pathcPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)
	if err := writeJSON(w, http.StatusOK, post); err != nil {
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

		ctx = context.WithValue(ctx, "post", post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value("post").(*store.Post)
	return post
}
