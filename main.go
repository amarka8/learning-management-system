package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "github.com/amarka8/learning-management-system/cmd/server"
	"github.com/amarka8/learning-management-system/config"
	"github.com/amarka8/learning-management-system/environment"
	logger "github.com/amarka8/learning-management-system/logger"
)

func main() {
	var server http.Server
	var port string

	// var err error
	// var schemafilename string
	// var tokenfilename string
	// var handleDbRequests http.Handler

	// TODO: add debug/verbose flag
	flag.StringVar(&port, "p", "8080", "port number")
	flag.Parse()

	logging := logger.NewRealLogger()
	env := environment.Environment(logging, config.NullConfig{})

	/*
		register the port which server is listening on along with handler which
		spawns new goroutines when client sends request to server
	*/
	server = http.Server{Addr: ":" + port}
	server.Handler = api.NewDataDocHandler(env)

	/*
		Ensures that we can ctrl+c to kill the application while it runs on the command line
	*/
	// signal.Notify requires the channel to be buffered
	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)
	go func() {
		// Wait for Ctrl-C signal
		<-ctrlc
		logging.Info("Server shutting down")
		server.Close()
	}()

	logging.Info("Serving on port 8080 at localhost")
	server.ListenAndServe()
}
