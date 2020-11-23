package api

import (
	response "feedback-service/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"feedback-service/pkg/storage"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "server started")
}

func handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var rq storage.ReviewQuery

	q := r.URL.Query()
	rq.Sort = q.Get("sort")
	rq.Answers = q.Get("answers")
	rq.Page, _ = strconv.Atoi(q.Get("page"))

	reviews, err := storage.GetReviews(rq)
	if err != nil {
		response.New("select error", false).WithError(err).Respond(w)
	}
	resp := storage.GetResponse(storage.GetPagesCount(), rq.Page, reviews)

	response.New("success", true).WithData(&resp).Respond(w)
}
