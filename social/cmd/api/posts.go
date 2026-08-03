package main

import (
	"context"
	"errors"
	"main/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

type postKey string

const postCtx postKey = "post"

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	if err := Validate.Struct(&payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return

	}
	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		//TODO: Change after auth
		UserID: 1,
	}
	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)
	comments, err := app.store.Comments.GetPostByID(r.Context(), post.ID)
	if err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}
	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}

}
func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}
	ctx := r.Context()
	if err := app.store.Posts.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.statusInternalServerErrorHandler(w, r, err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdatePostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	if err := Validate.Struct(&payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if err := app.store.Posts.Update(r.Context(), post); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}

}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.statusInternalServerErrorHandler(w, r, err)
			return
		}
		ctx := r.Context()
		post, err := app.store.Posts.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.statusInternalServerErrorHandler(w, r, err)
			}
			return
		}

		comments, err := app.store.Comments.GetPostByID(ctx, id)
		if err != nil {
			app.statusInternalServerErrorHandler(w, r, err)
			return
		}
		post.Comments = comments

		if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
			app.statusInternalServerErrorHandler(w, r, err)
			return
		}
		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value(postCtx).(*store.Post)
	return post
}

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
	UserID  int64  `json:"user_id" validate:"required"`
}

func (app *application) createCommentPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	if err := Validate.Struct(&payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return

	}

	post := getPostFromCtx(r)

	ctx := r.Context()
	comment := &store.Comment{
		Content: payload.Content,
		UserID:  payload.UserID,
		PostID:  post.ID,
	}
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}

}
