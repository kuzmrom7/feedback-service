package repository

import "time"

type Review struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	OrderHash string    `json:"order_hash"`
	Rated     string    `json:"rated" gorm:"type:time"`
	Rating    int       `json:"rating"`
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

type Pages struct {
	Total int `json:"total" db:"total"`
}

type ReviewsRepository interface {
	GetReviews(rq ReviewQuery) ([]Review, error)
	GetLastReview() (Review, error)
	GetReviewsCount() (int64, error)
	AddReviews(rw []Review) error
}