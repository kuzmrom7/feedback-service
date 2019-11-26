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

	rq.Sort = r.URL.Query().Get("sort")
	rq.Filter = r.URL.Query().Get("filter")

	fmt.Println(rq.Sort, "===")

	storage.GetList(rq).Respond(w)
}
