package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Single(d Data) Response {
	meta := Meta{Code: http.StatusOK, Message:"Successfully Executed"}
	return Response{Meta: meta, Data:d.Map()}
}

func List(d gin.H) Response {
	meta := Meta{Code: http.StatusOK, Message:"Successfully Executed"}
	return Response{Meta: meta, Data: d}
}
