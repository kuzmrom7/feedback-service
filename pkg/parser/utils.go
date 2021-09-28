package parser

import (
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
)

func getUrl(baseUrl string, chainId int64, limit int, offset int) string {
	return fmt.Sprintf("%s/reviews?chainId=%v&limit=%v&offset=%v&cacheBreaker=1572361583", baseUrl, chainId, limit, offset)
}

func mappedTypes(rw *Reviews, chainId int64) repository.Reviews {
	reviews := repository.Reviews{}
	reviews.Total = rw.Total

	for _, r := range rw.Reviews {
		review := repository.Review{
			Author:    r.Author,
			Body:      r.Body,
			OrderHash: r.OrderHash,
			RatedAt:   r.Rated,
			PlaceId:   chainId,
			Rate:      r.Icon,
		}

		var answers []repository.Answer

		for _, a := range r.Answers {
			answer := repository.Answer{
				Answer:    a.Answer,
				CreatedAt: a.CreatedAt,
				SourceId:  a.SourceId,
				StatusId:  a.StatusId,
			}

			answers = append(answers, answer)
		}

		review.Answers = answers

		reviews.Data = append(reviews.Data, review)
	}

	return reviews
}