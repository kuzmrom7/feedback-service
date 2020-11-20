package storage

import (
	response "feedback-service/pkg/utils"
	"log"
)

type Pages struct {
	Total int `json:"total" db:"total"`
}

func (r *Reviews) WriteMany() *response.Response {
	db.Create(&r.Data)
	return nil
}

func GetReviewCount() int64 {
	var total int64

	db.Model(&Review{}).Count(&total)
	return total
}

func GetReviews(rq ReviewQuery) *response.Response {
	var (
		reviews []Review
	)

	if _, err := db.Limit(10).Offset(0).Find(&reviews).DB(); err != nil {
		log.Println(err)
		return response.New("select error", false).WithError(err)
	}

	return response.New("success", true).WithData(&reviews)
}

func GetLast() (Review, error) {
	var review Review
	db.Order("rated desc").Find(&review)

	return review, nil
}
