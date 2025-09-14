package main

import (
	"net/http"
)

// Declare a handle which writes plain-text response with information about the application status, operating env and version
func (app *application) health(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds our response information
	envlp := envelope{
		"status": "available",
		"status_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envlp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
