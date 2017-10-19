package tools

type SearchRequest struct {
	Page int `json:"page" form:"page" validate:"-"`
	Size int `json:"size" form:"size" validate:"-"`
}

type Paginate struct {
	Size  int         `json:"size"`
	Page  int         `json:"page"`
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}
