package storage

import (
	response "feedback-service/pkg/utils"
	"log"
	"math"
)

const LIMIT = 100

type Pages struct {
	Total int `json:"total" db:"total"`
}

func (r *Reviews) WriteMany() *response.Response {
	db.Create(&r.Data)
	return nil
}

func GetPagesCount() int {
	var total int64

	db.Model(&Review{}).Count(&total)
	totalPages := int(math.Ceil(float64(total) / float64(LIMIT)))

	return totalPages
}

func GetReviews(rq ReviewQuery) ([]Review, error) {
	var (
		reviews []Review
		offset  int
	)

	if rq.Page > 0 {
		offset = (rq.Page * LIMIT) - 100
	}

	if _, err := db.Limit(LIMIT).Offset(offset).Order("rated desc").Find(&reviews).DB(); err != nil {
		log.Println(err)
		return nil, err
	}

	return reviews, nil
}

func GetLast() (Review, error) {
	var review Review
	db.Order("rated desc").Find(&review)

	return review, nil
}
