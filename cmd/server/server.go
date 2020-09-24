package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

// Start server
func Start(c *cli.Context) error {
	port := c.String("port")

	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
