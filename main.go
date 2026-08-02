package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "github.com/amarka8/learning-management-system/cmd/server"
	logger "github.com/amarka8/learning-management-system/logger"
)

func main() {
	var server http.Server
	var loginUser string
	var dbHost string
	var dbUser string
	var dbPassword string
	var url string
	var port string

	// var err error
	// var schemafilename string
	// var tokenfilename string
	// var handleDbRequests http.Handler

	// TODO: add debug/verbose flag
	flag.StringVar(&port, "p", "8080", "port number")
	flag.StringVar(&loginUser, "login-user", "", "login username")
	flag.StringVar(&dbHost, "host", "localhost", "database host (e.g. mysql, mariadb, etc.)")
	flag.StringVar(&dbUser, "db-user", "", "database username")
	flag.StringVar(&dbPassword, "db-password", "", "database password")
	flag.StringVar(&url, "url", "", "database URL")
	flag.Parse()

	logging := logger.NewRealLogger()

	/*
		register the port which server is listening on along with handler which
		spawns new goroutines when client sends request to server
	*/
	server = http.Server{Addr: ":" + port}
	server.Handler = api.NewDataDocHandler(dbHost, dbUser, dbPassword, logging)

	/*
		Ensures that we can ctrl+c to kill the application while it runs on the command line
	*/
	// signal.Notify requires the channel to be buffered
	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)
	go func() {
		// Wait for Ctrl-C signal
		<-ctrlc
		server.Close()
	}()

	logging.Info("Serving on port 8080 at localhost")
	server.ListenAndServe()
}
