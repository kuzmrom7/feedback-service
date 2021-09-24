package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"github.com/kuzmrom7/feedback-service/pkg/config"
)

func routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", Hello)
	r.GET("/reviews", handleGetReviews)

	return r
}

func Run(cfg config.Server) error {
	routes := routes()

	var handler http.Handler

	handler = handlers.LoggingHandler(os.Stdout, routes)

	p := strconv.Itoa(cfg.Port)

	srv := &http.Server{
		Addr:    ":" + p,
		Handler: handler,
	}

	log.Printf("server started on %s\n", p)
	return srv.ListenAndServe()
}
