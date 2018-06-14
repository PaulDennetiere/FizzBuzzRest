package main

import (
	"context"
	"fizzbuzz_rest/config"
	"fizzbuzz_rest/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	// Since this could be deployed anywhere, the logs are directed to STDOUT. This will be easily redirected when needed (eg: ./binary > myfile.log).
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)

	// Getting the configuration
	var serverConfig config.Server
	err := envconfig.Process("fizzbuzz", &serverConfig)
	if err != nil {
		errorLogger.Println("Could not parse configuration. Are your environment variables set correctly (see README.md) ? " + err.Error())
		infoLogger.Println("Exiting...")
		return
	}

	// Setting up the router.
	r := mux.NewRouter()
	recover := handlers.Recover{
		ErrorLogger: errorLogger,
	}
	r.Handle("/", recover.GenerateMiddleware(handlers.Fizzbuzz{})).Methods("GET")
	r.MethodNotAllowedHandler = handlers.BadMethod{}
	r.NotFoundHandler = handlers.NotFound{}

	// Setting up the server.
	srv := &http.Server{
		Addr:        "0.0.0.0:" + serverConfig.Port,
		ReadTimeout: serverConfig.Timeout,
		Handler:     r,
	}

	// Running the server.
	go func() {
		infoLogger.Println("Starting the server on: " + srv.Addr)
		infoLogger.Println("Configuration is read from env variables")
		if err := srv.ListenAndServe(); err != nil {
			errorLogger.Println(err.Error())
		}
	}()

	// Handling gracefull stops
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), serverConfig.GracefullStop)
	defer cancel()
	srv.Shutdown(ctx)
	infoLogger.Println("Gracefully stop the server.")
	os.Exit(0)
}
