package models

type (
	PaginateReq struct {
		Start string `query:"start" validate:"max=64"`
		Limit int    `query:"limit" validate:"required,min=2,max=100"`
	}

	PaginateRes struct {
		Data  any             `json:"data"`
		Limit int             `json:"limit"`
		Total int64           `json:"total"`
		First FirstPaginate   `json:"first"`
		Next  NextPaginateRes `json:"next,omitempty"`
	}

	FirstPaginate struct {
		Href  string `json:"href"`
		Start string `json:"start"`
	}

	NextPaginateRes struct {
		Href string `json:"href"`
	}

	KafkaOffset struct {
		Offset int64 `json:"offset" bson:"offset"`
	}
)
