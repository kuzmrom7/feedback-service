package parser

import (
	"encoding/json"
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

var (
	Token      string
	httpClient = &http.Client{}
	cooks      []*http.Cookie
	total      int
	lastReview repository.Review
	parsed     bool
)

const (
	baseUrl = "https://server.delivery-club.ru/api1.2"
	limit   = 1000
	chainId = 28720
)

type Parser struct {
	reviewsRepository repository.ReviewsRepository
}

func NewParser(rw repository.ReviewsRepository) *Parser {
	return &Parser{reviewsRepository: rw}
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
	url := fmt.Sprintf("%s/user/login", baseUrl)
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

	Token = fmt.Sprintf("%v", result["token"])
	cooks = resp.Cookies()
}

func (p *Parser) getReviews() {
	offset := 0

	p.request(0)

	steps := total / limit

	for i := 0; i < steps; i++ {
		if parsed {
			continue
		}
		offset = offset + limit
		p.request(offset)
		log.Println("Parsed", offset, "reviews")
	}
}

func (p *Parser) request(offset int) {
	url := fmt.Sprintf("%s/reviews?chainId=%v&limit=%v&offset=%v&cacheBreaker=1572361583", baseUrl, chainId, limit, offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	for _, cookie := range cooks {
		req.AddCookie(cookie)
	}

	req.Header.Set("x-user-authorization", Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	reviews := &repository.Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &reviews)
	if err != nil {
		log.Println(err)
	}

	/* Reviews that contain the latest in the repository */
	if !reflect.DeepEqual(repository.Reviews{}, lastReview) {
		data := sliceExtra(reviews.Data)
		if data == nil {
			return
		}

		reviews.Data = data
	}

	/* Save to repository */
	err = p.reviewsRepository.AddReviews(reviews.Data)
	if err != nil {
		log.Println(err)
		return
	}

	total = reviews.Total

}

func sliceExtra(reviews []repository.Review) []repository.Review {

	for i, review := range reviews {

		if review.OrderHash == lastReview.OrderHash {

			slicedRws := reviews[0:i]
			parsed = true

			if len(slicedRws) == 0 {
				return nil
			}

			log.Println("Detected", len(slicedRws), "new reviews!")

			return slicedRws
		}
	}

	return reviews
}
