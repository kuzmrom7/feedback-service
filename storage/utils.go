package storage

import (
	"fmt"
	"strconv"
)

func getListQueryBuilder(rq ReviewQuery) string {

	sort := "rated"
	limit := "100"
	offset := 0
	filters := ""

	if len(rq.Sort) > 0 {
		sort = rq.Sort
	}

	if len(rq.Answers) > 0 {
		if rq.Answers == "true" {
			filters = "where answers != '[]'"
		}
	}

	if rq.Page > 0 {
		l, _ := strconv.Atoi(limit)
		offset = (rq.Page * l) - 100
	}

	q := fmt.Sprintf(`SELECT *
		FROM review r
		%s
		ORDER BY %s DESC
		OFFSET %v
		limit %s`, filters, sort, offset, limit)

	return q

}

func getResponse(ps int, page int, reviews []Review) ResponseReviews {

	nxtp := 0

	if ps >= page+1 {
		nxtp = page + 1
	} else {
		nxtp = -1
	}

	resp := ResponseReviews{Pages: ps, Page: page, Reviews: reviews, NextPage: nxtp}

	return resp

}
