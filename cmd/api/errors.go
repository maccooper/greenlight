package main

import (
	"fmt"
	"net/http"
)

// logError(), generic helper
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse(), generic helper for JSON formatted eror messages to client
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	envlp := envelope{"error": message}

	err := app.writeJSON(w, status, envlp, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// errors param of type map[string]string, matches our Validator type.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}


// serverErrorResponse(), logs detailed errors through errorResponse() helper
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResposne(), sends 404 and json response to client
func (app *application) notFoundResposne(w http.ResponseWriter, r *http.Request) {
	message := "the required resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

//methodNotAllowedResponse(), sends 405 and json response to client
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for the resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
