package storage

import (
	response "feedback-service/utils"
	"fmt"
	"log"
)

func (r *Reviews) Write() *response.Response {

	_, err := db.NamedExec("INSERT INTO review ( author,answers, body, orderhash, rated,rating) VALUES ( :author, :answers, :body, :orderhash, :rated, :rating)", r.Data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func queryBuilder(rq ReviewQuery) string {

	sort := "rated"
	limit := "100"
	offset := "0"

	if len(rq.Sort) > 0 {
		sort = rq.Sort
	}

	if len(rq.Limit) > 0 {
		limit = rq.Limit
	}

	if len(rq.Offset) > 0 {
		offset = rq.Offset
	}

	q := fmt.Sprintf(`SELECT *
		FROM review r
		ORDER BY %s DESC
		OFFSET %s
		limit %s`, sort, offset, limit)

	return q

}

func GetList(rq ReviewQuery) *response.Response {

	var (
		reviews Reviews
		q       string
	)

	q = queryBuilder(rq)

	if err := db.Select(&reviews.Data, q); err != nil {
		log.Println(err)
		return response.New("select error", false).WithError(err)
	}

	if len(reviews.Data) == 0 {
		return response.New("not found any products", true)
	}
	return response.New("success", true).WithData(reviews.Data)
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
