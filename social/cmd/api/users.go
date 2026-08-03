package main

import (
	"context"
	_ "errors"
	"main/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

// GetUser godoc
//
// @Summary		Fetches a user profile
// @Description	Fetches a user profile by ID
// @Tags 		users
// @Param		id	path	int	true	"User ID"
// @Success		200	{object}	store.User
// @Failure		400	{object}	error
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Security	ApiKeyAuth
// @Router		/users/{id}	[get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromContext(r)
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := app.getUserFromContext(r)
	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
		return
	}
	var payload FollowUser
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	ctx := r.Context()
	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrorConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.statusInternalServerErrorHandler(w, r, err)
		}

	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := app.getUserFromContext(r)
	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
	}
	var payload FollowUser
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	ctx := r.Context()
	if err := app.store.Followers.Unfollow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrorConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.statusInternalServerErrorHandler(w, r, err)
		}
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestResponsee(w, r, err)
			return
		}
		ctx := r.Context()
		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			switch err {
			case store.ErrNotFound:
				app.notFoundResponse(w, r, err)
				return
			default:
				app.statusInternalServerErrorHandler(w, r, err)
				return
			}
		}

		ctx = context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (app *application) getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
