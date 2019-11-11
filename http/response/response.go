package response

type Response struct {
	Meta Meta  `json:"meta"`
	Data interface{}      `json:"data, omitempty"`
}

type Full struct {
	MetaFull MetaFull `json:"meta"`
	PageInfo PageInfo `json:"page_info"`
	Errors   Errors   `json:"errors"`
	Data     DataFull `json:"data, omitempty"`
}