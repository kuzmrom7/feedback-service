package api

import (
	"feedback-service/storage"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "server started")
}

func handleGetReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var rq storage.ReviewQuery

	q := r.URL.Query()
	rq.Sort = q.Get("sort")
	rq.Filter = q.Get("filter")
	rq.Offset = q.Get("offset")
	rq.Limit = q.Get("limit")

	storage.GetList(rq).Respond(w)
}
