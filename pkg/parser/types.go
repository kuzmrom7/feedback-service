package parser

import (
	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"net/http"
)

type Parser struct {
	cfg               config.Parser
	reviewsRepository repository.ReviewsRepository
	token             string
	cooks             []*http.Cookie
	total             int
	parsed            bool
	lastReview        repository.Review
}

type Reviews struct {
	Reviews []Review `json:"reviews"`
	Total   int      `json:"total"`
}

type Review struct {
	Answers   []Answer  `json:"answers"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	Icon      string    `json:"icon"`
	OrderHash string    `json:"orderHash"`
	Rated     string    `json:"rated"`
	Products  []Product `json:"products"`
}

type Product struct {
	Name string `json:"name"`
}

type Answer struct {
	Answer    string `json:"answer"`
	CreatedAt string `json:"createdAt"`
	SourceId  string `json:"sourceId"`
	StatusId  string `json:"statusId"`
}
