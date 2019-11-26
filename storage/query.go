package storage

import "fmt"

func getListQueryBuilder(rq ReviewQuery) string {

	sort := "rated"
	limit := "100"
	offset := "0"
	filters := ""

	if len(rq.Limit) > 0 {
		limit = rq.Limit
	}

	if len(rq.Offset) > 0 {
		offset = rq.Offset
	}

	if len(rq.Answers) > 0 {
		if rq.Answers == "true" {
			filters = "where answers != '[]'"
		}
	}

	q := fmt.Sprintf(`SELECT *
		FROM review r
		%s
		ORDER BY %s DESC
		OFFSET %s
		limit %s`, filters, sort, offset, limit)

	return q

}
