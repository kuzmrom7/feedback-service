package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"math"
)

type ResponseReviews struct {
	Reviews  []repository.Review `json:"reviews"`
	Total    int                 `json:"totalPages"`
	Page     int                 `json:"page"`
	NextPage int                 `json:"nextPage"`
}

func NewResponseReviews(reviewsTotal int64, page int, reviews []repository.Review) *ResponseReviews {
	pagesCount := int(math.Ceil(float64(reviewsTotal) / float64(LIMIT)))

	resp := &ResponseReviews{Reviews: reviews, Total: pagesCount, Page: page}

	if page == 0 {
		resp.Page = 1
	}

	if page > 0 {
		resp.NextPage = page + 1
	}

	if page >= pagesCount {
		resp.NextPage = 1
	}

	return resp
}

type ResponseError struct {
	Message string `json:"message"`
}
