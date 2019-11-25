package parser

import (
	"encoding/json"
	"feedback-service/storage"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Token      string
	httpClient = &http.Client{}
	cooks      []*http.Cookie
	total      int
)

const (
	baseUrl = "https://api.delivery-club.ru/api1.2"
	limit   = 1000
	chainId = 28720
)

func Run() {
	getToken()
	getReviews()

	//storage.GetOne()
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
		offset = offset + limit
		request(offset)
		fmt.Println(steps)
	}

}

func request(offset int) {

	url := fmt.Sprintf("%s/reviews?chainId=%v&limit=%v&offset=%v&cacheBreaker=1572361583", baseUrl, chainId, limit, offset)

	log.Println("offset = ", offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, cookie := range cooks {
		req.AddCookie(cookie)
	}

	req.Header.Set("x-user-authorization", Token)
	fmt.Println(Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	reviews := &storage.Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(body, &reviews)

	reviews.Write()

	total = reviews.Total

}
