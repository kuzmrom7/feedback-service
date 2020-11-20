package storage

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Author    string `json:"author"`
	Body      string `json:"body"`
	OrderHash string `json:"orderHash"`
	Rated     string `json:"rated" gorm:"type:time"`
	Rating    int    `json:"rating"`
}

type Reviews struct {
	Data  []Review `json:"reviews"`
	Total int      `json:"total"`
}

type ReviewQuery struct {
	Sort    string
	Answers string
	Page    int
}

type ResponseReviews struct {
	Reviews  []Review `json:"reviews"`
	Pages    int      `json:"totalPages"`
	Page     int      `json:"currentPage"`
	NextPage int      `json:"nextPage"`
}
