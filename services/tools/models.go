package tools

type SearchRequest struct {
	Start  int64             `json:"start" validate:"-"`
	Length int64             `json:"length" validate:"-"`
	Query  map[string]string `json:"query" validate:"-"`
}

type Paginate struct {
	Length int64       `json:"length"`
	Offset int64       `json:"offset"`
	Total  int64       `json:"total"`
	Items  interface{} `json:"items"`
}
