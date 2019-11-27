package api

import (
	"feedback-service/storage"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "server started")
}

func handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var rq storage.ReviewQuery

	q := r.URL.Query()
	rq.Sort = q.Get("sort")
	rq.Limit = q.Get("limit")
	rq.Answers = q.Get("answers")
	rq.Page, _ = strconv.Atoi(q.Get("page"))

	storage.GetList(rq).Respond(w)
}
