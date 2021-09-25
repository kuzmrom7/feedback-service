package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository/postgres"
	response "github.com/kuzmrom7/feedback-service/pkg/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var rq postgres.ReviewQuery

	q := r.URL.Query()
	rq.Sort = q.Get("sort")
	rq.Answers = q.Get("answers")
	rq.Page, _ = strconv.Atoi(q.Get("page"))

	reviews, err := postgres.GetReviews(rq)
	if err != nil {
		response.New("select error", false).WithError(err).Respond(w)
	}
	resp := createResponse(postgres.GetPagesCount(), rq.Page, reviews)

	response.New("success", true).WithData(&resp).Respond(w)
}


func createResponse(pagesCount int, page int, reviews []postgres.Review) ResponseReviews {
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