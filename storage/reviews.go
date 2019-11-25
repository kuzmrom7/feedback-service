package storage

import (
	response "feedback-service/utils"
	"log"
)



func (r *Reviews) Write() *response.Response {

	_, err := db.NamedExec("INSERT INTO review ( author, body, orderhash, rated,rating) VALUES ( :author, :body, :orderHash, :rated, :rating)", r.Data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//func GetOne() *response.Response {
//
//	r := Review{}
//
//	rows, err := db.NamedQuery(`SELECT * FROM review ORDER BY rated desc limit 1`, r)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(rows, "Hi &&")
//
//	return nil
//
//}
