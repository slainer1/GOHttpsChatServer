package main

import (
	"main/internal/store"
	"net/http"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required, max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// registerUserHandler godoc
//
//	@Summary Registers a User
//	@Description Registers a user
//	@Tags		authentication
//	@Accept		json
//	@Produce	json
//	@param		payload body	RegisterUserPayload true	"User credentials"
//	@Success		201		{object}	store.User		"User registered"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponsee(w, r, err)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}
	//hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
	}
	//store the user
	err := app.store.Users.

	//success code
	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.statusInternalServerErrorHandler(w, r, err)
	}

}
