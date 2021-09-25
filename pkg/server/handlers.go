package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	response "github.com/kuzmrom7/feedback-service/pkg/utils"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// LIMIT TODO: update limit
const LIMIT = 100

func (s *Server) handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var rq repository.ReviewQuery

	q := r.URL.Query()
	rq.Sort = q.Get("sort")
	rq.Answers = q.Get("answers")
	rq.Page, _ = strconv.Atoi(q.Get("page"))

	reviews, err := s.reviewsRepository.GetReviews(rq)
	if err != nil {
		response.New("select error", false).WithError(err).Respond(w)
	}

	reviewTotal, err := s.reviewsRepository.GetReviewsCount()
	if err != nil {
		return
	}
	resp := s.createResponse(reviewTotal, rq.Page, reviews)

	response.New("success", true).WithData(&resp).Respond(w)
}

//TODO:fix this method
func (s *Server) createResponse(reviewsTotal int64, page int, reviews []repository.Review) ResponseReviews {
	pagesCount := int(math.Ceil(float64(reviewsTotal) / float64(LIMIT)))
	if page == 0 {
		page = 1
	}

	nextPage := 0

	if page <= 0 {
		nextPage = 0
		page = 0
	} else if pagesCount >= page+1 {
		nextPage = page + 1
	} else {
		nextPage = -1
	}

	resp := ResponseReviews{Pages: pagesCount, Page: page, Reviews: reviews, NextPage: nextPage}

	return resp
}
