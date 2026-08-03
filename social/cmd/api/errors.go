package main

import (
	"log"
	"net/http"
)

func (app *application) statusInternalServerErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "The server encountered an problem")
}

func (app *application) badRequestResponsee(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, "resource not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflict Error: %s path: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusConflict, "resource not found")
}
