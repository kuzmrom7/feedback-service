package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kuzmrom7/feedback-service/pkg/config"
)

type Server struct {
	cfg               config.Server
	reviewsRepository repository.ReviewsRepository
}

func New(cfg config.Server, r repository.ReviewsRepository) *Server {
	return &Server{cfg: cfg, reviewsRepository: r}
}

func (s *Server) Run() error {
	router := httprouter.New()

	router.GET("/reviews", s.handleGetReviews)

	p := strconv.Itoa(s.cfg.Port)

	srv := &http.Server{
		Addr:    ":" + p,
		Handler: router,
	}

	log.Printf("server started on %s\n", p)
	return srv.ListenAndServe()
}
