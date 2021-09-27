package parser

import (
	"encoding/json"
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"log"
	"net/http"
	"reflect"
)

var (
	token      string
	httpClient = &http.Client{}
	cooks      []*http.Cookie
	total      int
	lastReview repository.Review
	parsed     bool
)

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

func (p *Parser) setToken() {
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

	token = fmt.Sprintf("%v", result["token"])
	cooks = resp.Cookies()
}

func (p *Parser) getReviews() {
	offset := 0

	url := getUrl(p.cfg.BaseURL, p.cfg.ChainId, p.cfg.Limit, offset)
	reviews := requestReviews(url)

	p.saveReviews(reviews)

	steps := total / p.cfg.Limit

	for i := 0; i < steps; i++ {
		if parsed {
			break
		}

		offset = offset + p.cfg.Limit

		url = getUrl(p.cfg.BaseURL, p.cfg.ChainId, p.cfg.Limit, offset)
		reviews := requestReviews(url)

		p.saveReviews(reviews)

		log.Println("Parsed", offset, "reviews")
	}
}

func (p *Parser) saveReviews(rw *Reviews) {
	reviews := mappedTypes(rw, p.cfg.ChainId)

	/* Reviews that contain the latest in the repository */
	if !reflect.DeepEqual(repository.Reviews{}, lastReview) {
		data := sliceExtra(reviews.Data)
		if data == nil {
			return
		}

		reviews.Data = data
	}

	/* Save to repository */
	err := p.reviewsRepository.AddReviews(reviews.Data)
	if err != nil {
		log.Println(err)
		return
	}

	total = reviews.Total
}
