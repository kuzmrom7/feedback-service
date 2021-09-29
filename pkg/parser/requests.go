package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

func (p *Parser) requestReviews() *Reviews {
	url := getUrl(p.cfg.BaseURL, p.cfg.ChainId, p.cfg.Limit, p.offset)

	if p.lastFound {
		time.Sleep(1 * time.Second)
	}

	log.Println("New request to reviews url", url)
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
