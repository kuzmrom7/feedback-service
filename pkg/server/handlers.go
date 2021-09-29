package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const LIMIT = 100

func (s *Server) handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	var rq repository.ReviewQuery
	respErr := &ResponseError{}

	q := r.URL.Query()

	qPage := q["page"]
	qSort := q.Get("sort")

	if qPage == nil {
		qPage = append(qPage, "1")
	}

	page, err := strconv.Atoi(qPage[0])

	if err != nil {
		respErr.Message = err.Error()
		respErr.Respond(w, http.StatusInternalServerError)
		return
	}

	rq.Page = page
	rq.Sort = qSort

	reviews, err := s.reviewsRepository.GetReviews(rq)
	if err != nil {
		respErr.Message = err.Error()
		respErr.Respond(w, http.StatusInternalServerError)
		return
	}

	reviewTotal, err := s.reviewsRepository.GetReviewsCount()
	if err != nil {
		respErr.Message = err.Error()
		respErr.Respond(w, http.StatusInternalServerError)
		return
	}

	NewResponseReviews(reviewTotal, rq.Page, reviews).SuccessRespond(w)
}
