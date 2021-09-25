package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kuzmrom7/feedback-service/pkg/config"
)

type Server struct {
	cfg config.Server
}

func New(cfg config.Server) Server {
	return Server{cfg: cfg}
}

func (s *Server) Run() error {
	router := httprouter.New()

	router.GET("/reviews", handleGetReviews)

	p := strconv.Itoa(s.cfg.Port)

	srv := &http.Server{
		Addr:    ":" + p,
		Handler: router,
	}

	log.Printf("server started on %s\n", p)
	return srv.ListenAndServe()
}
