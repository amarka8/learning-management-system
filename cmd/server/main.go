// This Package is responsible for defining the API requests to create new datadocs, which represent new
// learning journeys/timelines

package server

import (
	"encoding/json"
	"net/http"

	"github.com/amarka8/learning-management-system/environment"
)

type Auth interface {
	authorize(tok string) bool
}

type querybook struct {
	env *environment.Env
}

type healthcheck struct {
	Status string
}

func NewDataDocHandler(env *environment.Env) http.Handler {
	if env == nil {
		env = environment.Null()
	}

	qbook := querybook{env: env}
	mux := http.NewServeMux()
	mux.HandleFunc("OPTIONS /health", optionsHealthCheck)
	mux.HandleFunc("GET /health", qbook.health)
	qbook.env.Info("Constructed API Handler in server package")
	return mux
}

// get handles GET requests for /health
// It takes in a http.ResponseWriter and *http.Request and is an *querybook method.
// If successful, returns 200, otherwise error codes of 4xx or 5xx.
func (qb *querybook) health(w http.ResponseWriter, r *http.Request) {
	// this means that we can allow all origins to access our API. this is because its public
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// this means the server sends back a JSON response
	w.Header().Set("Content-Type", "application/json")
	encoded, err := json.Marshal(healthcheck{Status: "ok"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		qb.env.Error("Issue with Health Check")
		w.Write([]byte(`"` + err.Error() + `"`))
	}
	qb.env.Info("Completed Health Check")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

// this method handles the preflight requests a browser/http client send to our server to ensure the request can be handled
// it is used to prevent large requests that will surely fail from being sent
func optionsHealthCheck(w http.ResponseWriter, r *http.Request) {
	// these are the headers a browser/HTTP client is allowed to send when making a request to this server
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, accept, Authorization")
	// this means that we can allow all origins to access our API. this is because its public
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// we allow all origins to access the /health resource using GET requests
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusOK)
}
