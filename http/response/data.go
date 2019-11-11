package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-gin/utils"
)

func Single(d Data) Response {
	meta := Meta{Code: http.StatusOK, Message:"Successfully Executed"}
	return Response{Meta: meta, Data:d.Map()}
}

func List(d gin.H) Response {
	meta := Meta{Code: http.StatusOK, Message:"Successfully Executed"}
	return Response{Meta: meta, Data: d}
}

func SingleFullResponse(httpMethod string, d Data) (response Full) {
	meta := MetaFullComposer(http.StatusOK, utils.ApiVersion(), httpMethod, utils.OperationSuccessfullyExecuted)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(utils.OK, d.Map(), nil)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data:dataFull}
}


func ListFullResponse(httpMethod string, d gin.H) (response Full) {
	meta := MetaFullComposer(http.StatusOK, utils.ApiVersion(), httpMethod, utils.OperationSuccessfullyExecuted)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(utils.Success, nil, d)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data:dataFull}
}

func ListAndSingleFullResponse(httpMethod string, single Data ,list gin.H) (response Full) {
	meta := MetaFullComposer(http.StatusOK, utils.ApiVersion(), httpMethod, utils.OperationSuccessfullyExecuted)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(utils.Success, single, list)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func SingleCustomFullResponse(httpStatus int, httpMethod string, message string, d Data) (response Full) {
	meta := MetaFullComposer(httpStatus, utils.ApiVersion(), httpMethod, message)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(message, d.Map(), nil)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func ListCustomFullResponse(httpStatus int, httpMethod string, message string, d gin.H) (response Full) {
	meta := MetaFullComposer(httpStatus, utils.ApiVersion(), httpMethod, message)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(message, nil, d)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func ListAndSingleCustomFullResponse(
	httpStatus int,
	httpMethod string,
	message string,
	single Data,
	list gin.H) (response Full) {

	meta := MetaFullComposer(httpStatus, utils.ApiVersion(), httpMethod, message)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer("0")
	dataFull := DataFullJsonComposer(message, single, list)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}
