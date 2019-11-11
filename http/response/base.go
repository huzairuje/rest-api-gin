package response

import (
	"net/http"
	"rest-api-gin/utils"
)

func Success() (response Response) {
	meta := MetaComposer(http.StatusOK, "OK")
	return Response{Meta: meta}
}

func SuccessCreated(d Data) (response Response) {
	meta := MetaComposer(http.StatusOK, "Success Created")
	return Response{Meta: meta, Data:d.Map()}
}

func SuccessUpdated(d Data) (response Response) {
	meta := MetaComposer(http.StatusOK, "Success Updated")
	return Response{Meta: meta, Data:d.Map()}
}

func SuccessDeleted() (response Response) {
	meta := MetaComposer(http.StatusOK, "Success Deleted")
	return Response{Meta: meta}
}

func NotFound() (response Response) {
	meta := MetaComposer(http.StatusNotFound, "Not Found")
	return Response{Meta: meta}
}

func SuccessFullResponse(httpMethod string) (response Full) {
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
	dataFull := DataFullJsonComposer(utils.Success, nil, nil)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func ValidationFullResponse(httpMethod string, error string) (response Full) {
	meta := MetaFullComposer(http.StatusUnprocessableEntity, utils.ApiVersion(), httpMethod, utils.SomethingWentWrong)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer(error)
	dataFull := DataFullJsonComposer(utils.ValidationError, nil, nil)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func InternalServerErrorFullResponse(httpMethod string, error string) (response Full) {
	meta := MetaFullComposer(http.StatusUnprocessableEntity, utils.ApiVersion(), httpMethod, utils.SomethingWentWrong)
	pageInfo := PageInfoComposer(
		1,
		1,
		1,
		0,
		"null",
		"null",
		1,
		0)
	errorBag := ErrorsComposer(error)
	dataFull := DataFullJsonComposer(utils.InternalServerError, nil, nil)
	return Full{MetaFull: meta, PageInfo:pageInfo, Errors: errorBag, Data: dataFull}
}

func CustomFullResponse(httpStatus int, httpMethod string, message string, error string) (response Full) {
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
	errorBag := ErrorsComposer(error)
	dataFull := DataFullJsonComposer(message, nil, nil)
	return Full{MetaFull: meta, PageInfo: pageInfo, Errors: errorBag, Data: dataFull}
}
