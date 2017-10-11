package tools

type SearchRequest struct {
	Page int `json:"page" form:"page" validate:"-"`
	Size int `json:"size" form:"size" validate:"-"`
}

type Paginate struct {
	Length int         `json:"length"`
	Offset int         `json:"offset"`
	Total  int64       `json:"total"`
	Items  interface{} `json:"items"`
}
