package parser

import (
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"reflect"
)

func (p *Parser) addReviews(rw *Reviews) {
	reviews := mappedTypes(rw, p.cfg.ChainId)

	/* Reviews that contain the latest in the repository */
	if !reflect.DeepEqual(repository.Reviews{}, p.lastReview) {
		data := p.sliceExtra(reviews.Data)
		if data != nil {
			/* Save to repository */
			err := p.reviewsRepository.AddReviews(data)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Add", len(reviews.Data), " new reviews")
		}
	}

	p.total = reviews.Total
}

func (p *Parser) updateReviews(rw *Reviews) {
	//log.Println("update")
}

func (p *Parser) sliceExtra(reviews []repository.Review) []repository.Review {
	for i, review := range reviews {
		if review.OrderHash == p.lastReview.OrderHash {
			slicedRws := reviews[0:i]
			p.parsed = true

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