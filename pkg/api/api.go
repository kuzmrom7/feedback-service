package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", Hello)
	r.GET("/reviews", handleGetReviews)

	return r
}

func Run() error {
	routes := routes()

	var handler http.Handler

	handler = handlers.LoggingHandler(os.Stdout, routes)

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: handler,
	}

	log.Printf("server started on %s\n", "8080")
	return srv.ListenAndServe()
}
