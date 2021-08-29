package dto

type Searching struct {
	Search       []Movie `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

type Movie struct {
	Title  string `Json:"Title"`
	Year   string `Json:"Year"`
	ImdbID string `Json:"imdbID"`
	Type   string `Json:"Type"`
	Poster string `Json:"Poster"`
}
