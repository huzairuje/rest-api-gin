package response

import (
	"net/http"
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
