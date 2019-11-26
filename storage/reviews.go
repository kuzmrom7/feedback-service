package storage

import (
	response "feedback-service/utils"
	"log"
)

func (r *Reviews) Write() *response.Response {

	_, err := db.NamedExec("INSERT INTO review ( author, body, orderhash, rated,rating) VALUES ( :author, :body, :orderhash, :rated, :rating)", r.Data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetList() *response.Response {

	var reviews Reviews

	if err := db.Select(&reviews.Data, `
		SELECT r.id,r.author, r.body, r.orderhash, r.rated, r.rating, r.created, r.updated
		FROM review r
		ORDER BY rated desc
	`); err != nil {
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
