package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare our hard-coded version constant (we'll generate this automatically later in the book)
const version = "1.0.0"

// Define a config struct to hold our application settings.
type config struct {
	port int
	env  string
}

// Define an application struct to hold dependancies for http handlers, helpers and middleware.. It will grow as we add more over time
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	//Read the value of the port and env command-line flags into the config struct,
	//Defaults of port 4000 and dev-environment if no flags set
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	//Initialize a new logger which writes messages prefixed with current date time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Declare an instance of our application struct, containing config struct and the log
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Declare a new servemux and add /v1/health route which dispatches requests to the healthcheckHandler method
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/health", app.health)

	//Declare http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//Start thte http server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
