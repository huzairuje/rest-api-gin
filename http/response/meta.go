package response

import "github.com/gin-gonic/gin"

type Meta struct {
	Code		int 	`json:"code"`
	Message 	string 	`json:"message"`
}

func MetaComposer(code int, message string) Meta {
	meta := Meta{
		Code:    code,
		Message: message,
	}
	return meta
}

type MetaFull struct {
	Code       int    `json:"code"`
	ApiVersion string `json:"api_version"`
	Method     string `json:"method"`
	Message    string `json:"message"`
}

func MetaFullComposer(code int, apiVersion string, method string, message string) MetaFull {
	meta := MetaFull{
		Code:       code,
		ApiVersion: apiVersion,
		Method:     method,
		Message:    message,
	}
	return meta
}

type PageInfo struct {
	Total       int    `json:"total"`
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
	From        int    `json:"from"`
	To          int    `json:"to"`
}

func PageInfoComposer(
	total int,
	perPage int,
	currentPage int,
	lastPage int,
	nextPageUrl string,
	prevPageUrl string,
	from int,
	to int) PageInfo {

	pageInfo := PageInfo{
		Total:       total,
		PerPage:     perPage,
		CurrentPage: currentPage,
		LastPage:    lastPage,
		NextPageUrl: nextPageUrl,
		PrevPageUrl: prevPageUrl,
		From:        from,
		To:          to,
	}
	return pageInfo
}

type Errors string

func ErrorsComposer(message string) Errors {
	errors := Errors(message)
	return errors
}

type DataFull struct {
	Message string      `json:"message"`
	Item    interface{} `json:"item"`
	Items   gin.H    `json:"items"`
}

func DataFullJsonComposer(message string, item interface{}, items gin.H) DataFull {
	dataFull := DataFull{
		Message: message,
		Item:    item,
		Items:   items,
	}
	return dataFull
}