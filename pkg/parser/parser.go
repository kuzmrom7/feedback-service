package parser

import (
	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"math"
	"net/http"
)

var (
	httpClient = &http.Client{}
)

func NewParser(cfg config.Parser, rw repository.ReviewsRepository) *Parser {
	return &Parser{reviewsRepository: rw, cfg: cfg}
}

func (p *Parser) Run() {
	offset := 0

	review, err := p.reviewsRepository.GetLastReview()
	if err != nil {
		log.Println("Parser stop with errors = ",err)
		return
	}
	p.lastReview = review

	log.Println("Last review found = ", p.lastReview)
	p.setupCookies()

	// first req for getting totals
	reviews := p.requestReviews(offset)
	p.addReviews(reviews)

	steps := int(math.Ceil(float64(p.total) / float64(p.cfg.Limit)))

	for i := 0; i < steps; i++ {
		offset += p.cfg.Limit

		reviews := p.requestReviews(offset)

		log.Println("Parsed new", offset, "reviews")

		if p.parsed {
			p.updateReviews(reviews)
			continue
		}

		p.addReviews(reviews)
	}
}