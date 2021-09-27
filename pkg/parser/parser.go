package parser

import (
	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"net/http"
)

var (
	token      string
	httpClient = &http.Client{}
	cooks      []*http.Cookie
	total      int
	lastReview repository.Review
	parsed     bool
)

type Parser struct {
	cfg               config.Parser
	reviewsRepository repository.ReviewsRepository
}

func NewParser(cfg config.Parser, rw repository.ReviewsRepository) *Parser {
	return &Parser{reviewsRepository: rw, cfg: cfg}
}

func (p *Parser) Run() {
	review, err := p.reviewsRepository.GetLastReview()
	if err != nil {
		log.Println(err)
		return
	}

	lastReview = review

	log.Println("Last review found", lastReview)

	p.setToken()
	p.getReviews()
}
