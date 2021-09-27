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

	p.request(0)

	steps := total / p.cfg.Limit

	for i := 0; i < steps; i++ {
		if parsed {
			continue
		}
		offset = offset + p.cfg.Limit
		p.request(offset)
		log.Println("Parsed", offset, "reviews")
	}
}

func (p *Parser) request(offset int) {
	url := fmt.Sprintf("%s/reviews?chainId=%v&limit=%v&offset=%v&cacheBreaker=1572361583", p.cfg.BaseURL, p.cfg.ChainId, p.cfg.Limit, offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	for _, cookie := range cooks {
		req.AddCookie(cookie)
	}

	req.Header.Set("x-user-authorization", token)

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