package storage

import (
	response "feedback-service/utils"
	"log"
	"math"
)

func (r *Reviews) Write() *response.Response {

	_, err := db.NamedExec("INSERT INTO review ( author,answers, body, orderhash, rated,rating) VALUES ( :author, :answers, :body, :orderhash, :rated, :rating)", r.Data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

type Pages struct {
	Total int `json:"total" db:"total"`
}

func GetList(rq ReviewQuery) *response.Response {

	var (
		reviews Reviews
		q       string
	)

	var pages []Pages

	err := db.Select(&pages, `select count(*) as total from review`)
	if err != nil {
		log.Println(err)
		return response.New("select error", false).WithError(err)
	}

	p := float64(pages[0].Total) / float64(100)
	ps := int(math.Ceil(p))

	if rq.Page == 0 {
		rq.Page = 1
	}

	q = getListQueryBuilder(rq)

	if err := db.Select(&reviews.Data, q); err != nil {
		log.Println(err)
		return response.New("select error", false).WithError(err)
	}

	if len(reviews.Data) == 0 {
		return response.New("not found any reviews", true)
	}

	resp := getResponse(ps, rq.Page, reviews.Data)

	return response.New("success", true).WithData(resp)
}

func GetLast() ([]Review, error) {

	var reviews []Review

	if err := db.Select(&reviews, `
		SELECT * FROM review ORDER BY rated desc limit 1
		`); err != nil {
		log.Println(err)
		return reviews, err
	}

	return reviews, nil
}
