package postgres

import (
	"errors"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"gorm.io/gorm"
	"log"
)

const LIMIT = 100

type ReviewsRepository struct {
	db *gorm.DB
}

func NewReviewsRepository(db *gorm.DB) *ReviewsRepository {
	return &ReviewsRepository{db: db}
}

func (r *ReviewsRepository) AddReviews(rw []repository.Review) error {
	result := r.db.Create(rw)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ReviewsRepository) GetReviewsCount() (int64, error) {
	var total int64

	result := r.db.Model(&repository.Review{}).Count(&total)

	if result.Error != nil {
		return 0, result.Error
	}

	return total, nil
}

func (r *ReviewsRepository) GetReviews(rq repository.ReviewQuery) ([]repository.Review, error) {
	var (
		reviews []repository.Review
		offset  int
	)

	if rq.Page < 0 {
		return nil, errors.New("page cannot be negative number")
	}

	if rq.Page > 0 {
		offset = (rq.Page * LIMIT) - 100
	}

	if _, err := r.db.Limit(LIMIT).Offset(offset).Preload("Answers").Order("rated_at desc").Find(&reviews).DB(); err != nil {
		log.Println(err)
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewsRepository) GetLastReview() (repository.Review, error) {
	var review repository.Review
	result := r.db.Order("rated_at desc").Find(&review)

	if result.Error != nil {
		return review, result.Error
	}

	return review, nil
}

func (r *ReviewsRepository) UpdateReview(rw repository.Review) error {
	//result := r.db.Model(&repository.Review{}).
	//	Where(repository.Review{OrderHash: rw.OrderHash}).
	//	Assign(&rw).
	//	FirstOrCreate(&rw)

	//result:= r.db.Session(&gorm.Session{FullSaveAssociations: true}).
	//	Where(repository.Review{Body: rw.Body}).
	//	Updates(&rw)
	//
	//
	//if result.Error != nil {
	//	return result.Error
	//}

	return nil
}
