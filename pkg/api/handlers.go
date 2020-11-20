package api

import (
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

	storage.GetReviews(rq).Respond(w)
}
