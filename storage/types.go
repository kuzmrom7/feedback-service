package storage

import "encoding/json"

type Review struct {
	ID        string          `json:"id,omitempty" db:"id"`
	Author    string          `json:"author" db:"author"`
	Answers   json.RawMessage `json:"answers" db:"answers"`
	Body      string          `json:"body" db:"body"`
	OrderHash string          `json:"orderHash" db:"orderhash"`
	Rated     string          `json:"rated" db:"rated"`
	Rating    int             `json:"rating" db:"rating"`
	Created   string          `json:"created,omitempty" db:"created"`
	Updated   string          `json:"updated,omitempty" db:"updated"`
}

type Reviews struct {
	Data  []Review `json:"reviews"`
	Total int      `json:"total"`
}

type ReviewQuery struct {
	Sort   string
	Filter string
	Limit  string
	Offset string
}
