package parser

import (
	"encoding/json"
	"feedback-service/storage"
	"fmt"
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
	lastReview storage.Review
	parsed     bool
)

const (
	baseUrl = "https://api.delivery-club.ru/api1.2"
	limit   = 1000
	chainId = 28720
)

func Run() {
	reviews, err := storage.GetLast()
	if err != nil {
		log.Println(err)
	} else {
		if len(reviews) != 0 {
			lastReview = reviews[0]
			log.Println("Last review by", lastReview.Author)
		}

		getToken()
		getReviews()

	}

}

func getToken() {
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

func getReviews() {
	offset := 0

	request(0)

	steps := total / limit

	for i := 0; i < steps; i++ {
		if parsed {
			continue
		}
		offset = offset + limit
		request(offset)
		log.Println("Parsed", offset, "reviews")
	}

}

func request(offset int) {

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

	reviews := &storage.Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &reviews)
	if err != nil {
		log.Println(err)
	}

	if !reflect.DeepEqual(storage.Reviews{}, lastReview) {
		data := validate(reviews.Data)
		if data == nil {
			return
		}

		reviews.Data = data
	}

	reviews.Write()

	total = reviews.Total

}

func validate(rws []storage.Review) []storage.Review {

	for i := range rws {
		rw := rws[i]
		if rw.OrderHash == lastReview.OrderHash {

			slicedRws := rws[0:i]
			parsed = true

			if len(slicedRws) == 0 {
				return nil
			}

			log.Println("Detected", len(slicedRws), "new reviews!")

			return slicedRws
		}
	}
	return rws

}
