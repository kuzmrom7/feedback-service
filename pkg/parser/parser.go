package parser

import (
	"encoding/json"
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
	"time"
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

func (p *Parser) setupCookies() {
	url := fmt.Sprintf("%s/user/login", p.cfg.BaseURL)

	resp, err := httpClient.Post(url, "", nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	p.token = fmt.Sprintf("%v", result["token"])
	p.cooks = resp.Cookies()
}

func (p *Parser) requestReviews(offset int) *Reviews {
	url := getUrl(p.cfg.BaseURL, p.cfg.ChainId, p.cfg.Limit, offset)

	if p.parsed {
		time.Sleep(3 * time.Second)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	for _, cookie := range p.cooks {
		req.AddCookie(cookie)
	}
	req.Header.Set("x-user-authorization", p.token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	reviews := &Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &reviews)
	if err != nil {
		log.Println(err)
	}

	return reviews
}

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