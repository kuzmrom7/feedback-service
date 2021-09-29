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
	return &Parser{reviewsRepository: rw, cfg: cfg, offset: 0}
}

func (p *Parser) Run() {
	log.Println("Parser started....")

	review, err := p.reviewsRepository.GetLastReview()
	if err != nil {
		log.Println("Parser stop with errors = ", err)
		return
	}
	p.lastReview = review

	log.Println("Last review found")

	p.setupCookies()
	p.addReviews()

	steps := int(math.Ceil(float64(p.total) / float64(p.cfg.Limit)))

	for i := 0; i < steps; i++ {
		p.offset = p.offset + p.cfg.Limit

		if !p.lastFound {
			p.addReviews()
			continue
		}

		//p.updateReviews()
	}

	log.Println("Parser done")
}

func (p *Parser) addReviews() {
	rw := p.requestReviews()
	reviews := mappedTypes(rw, p.cfg.ChainId)

	data := p.sliceExtra(reviews.Data)

	if data != nil {
		err := p.reviewsRepository.AddReviews(data)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Add", len(reviews.Data), "new reviews")
	}

	p.total = reviews.Total
}

func (p *Parser) updateReviews() {
	rw := p.requestReviews()

	reviews := mappedTypes(rw, p.cfg.ChainId)
	for _, r := range reviews.Data {
		if err := p.reviewsRepository.UpdateReview(r); err != nil {
			log.Println(err)
		}
	}

	log.Println("Update", len(reviews.Data), "reviews")
}

//TODO: refactor need
func (p *Parser) sliceExtra(reviews []repository.Review) []repository.Review {
	for i, review := range reviews {
		if review.OrderHash == p.lastReview.OrderHash {
			slicedRws := reviews[0:i]

			p.lastFound = true
			p.offset = p.offset - p.cfg.Limit

			if len(slicedRws) == 0 {
				log.Println("Last review actual!")
				return nil
			}
			log.Println("Detected", len(slicedRws), "new reviews!")
			return slicedRws
		}
	}

	return reviews
}
