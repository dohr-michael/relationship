package tools

type SearchRequest struct {
	Start  int             `json:"start" validate:"-"`
	Length int             `json:"length" validate:"-"`
	Query  map[string]string `json:"query" validate:"-"`
}

type Paginate struct {
	Length int       `json:"length"`
	Offset int       `json:"offset"`
	Total  int       `json:"total"`
	Items  interface{} `json:"items"`
}
