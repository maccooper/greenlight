package main

import (
	"fmt"
	"net/http"
)

//Declare a handle which writes plain-text response with information about the application status, operating env and version
func (app *application) health (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
