package server

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
)

type ResponseReviews struct {
	Reviews  []repository.Review `json:"reviews"`
	Pages    int                 `json:"totalPages"`
	Page     int                 `json:"currentPage"`
	NextPage int                 `json:"nextPage"`
}
