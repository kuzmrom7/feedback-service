package storage

func GetResponse(pagesCount int, page int, reviews []Review) ResponseReviews {

	if page == 0 {
		page = 1
	}

	nextPage := 0

	if page <= 0 {
		nextPage = 0
		page = 0
	} else if pagesCount >= page+1 {
		nextPage = page + 1
	} else {
		nextPage = -1
	}

	resp := ResponseReviews{Pages: pagesCount, Page: page, Reviews: reviews, NextPage: nextPage}

	return resp
}
