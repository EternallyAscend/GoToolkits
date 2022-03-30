package response

import (
	"encoding/json"
	"fmt"
	gtkHttp "github.com/EternallyAscend/GoToolkit/pkg/network/https"
	gtkHeader "github.com/EternallyAscend/GoToolkit/pkg/network/https/header"
	"net/http"
)

type StringResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func makeResponseOK() StringResponse {
	return StringResponse{
		Status:  HttpResponseTrue,
		Message: gtkHttp.OKString,
		Code:    http.StatusOK,
	}
}

func (res StringResponse) Send(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(res)
	if nil != err {
		fmt.Println(err)
		return
	}
	gtkHeader.SetContentJsonHeader(w)
	w.WriteHeader(res.Code)
	_, err = w.Write(result)
	if nil != err {
		fmt.Println(err)
	}
}

func SendStringResponse(w http.ResponseWriter, r *http.Request, response StringResponse) {
	result, err := json.Marshal(response)
	if nil != err {
		fmt.Println(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(response.Code)
	_, err = w.Write(result)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func SendResponseOK(w http.ResponseWriter, r *http.Request) {
	makeResponseOK().Send(w, r)
}

func SendResponseInternalError(w http.ResponseWriter, r *http.Request, err error, code int) {
	StringResponse{
		Status:  code,
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}.Send(w, r)
}
