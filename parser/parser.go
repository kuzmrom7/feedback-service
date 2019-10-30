package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Token      string
	httpClient = &http.Client{}
	cooks      []*http.Cookie
)

const (
	baseUrl = "https://api.delivery-club.ru/api1.2"
)

func Run() {
	getToken()
	getReviews()
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
	url := fmt.Sprintf("%s/reviews?chainId=28720&limit=20&offset=0&cacheBreaker=1572361583", baseUrl)

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

	reviews := &Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(body, &reviews)

	//arr := reviews.Review
	//
	//fmt.Println(arr[0])
}
