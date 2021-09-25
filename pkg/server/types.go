package server

import "github.com/kuzmrom7/feedback-service/pkg/repository/postgres"

type ResponseReviews struct {
	Reviews  []postgres.Review `json:"reviews"`
	Pages    int               `json:"totalPages"`
	Page     int               `json:"currentPage"`
	NextPage int               `json:"nextPage"`
}
