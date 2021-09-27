package parser

import (
	"encoding/json"
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"io/ioutil"
	"log"
	"net/http"
)

func getUrl(baseUrl string, chainId int64, limit int, offset int) string {
	return fmt.Sprintf("%s/reviews?chainId=%v&limit=%v&offset=%v&cacheBreaker=1572361583", baseUrl, chainId, limit, offset)
}

func requestReviews(url string) *Reviews {
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

	//reviews := &repository.Reviews{}

	reviews := &Reviews{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	//err = json.Unmarshal(body, &reviews)
	//if err != nil {
	//	log.Println(err)
	//}

	///////
	err = json.Unmarshal(body, &reviews)
	if err != nil {
		log.Println(err)
	}
	/////

	return reviews
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
