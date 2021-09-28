package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	OrderHash string    `json:"order_hash"`
	RatedAt   string    `json:"rated_at" gorm:"type:time"`
	PlaceId   int64     `json:"place_id"`
	Rate      string    `json:"rate"`
	Answers   []Answer  `json:"answers" gorm:"foreignKey:ReviewId"`
}

func (a *Review) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return
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
