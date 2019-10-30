package parser

type Review struct {
	Answers   interface{} `json:"answers"`
	Author    string      `json:"author"`
	Body      string      `json:"body"`
	OrderHash string      `json:"orderHash"`
	Rated     string      `json:"rated"`
	Rating    string      `json:"rating"`
}

type Reviews struct {
	Review []Review `json:"reviews"`
}
