package request

import (
	"errors"
	"github.com/EternallyAscend/GoToolkits/pkg/network/https/header"
	"github.com/EternallyAscend/GoToolkits/pkg/network/https/response"
	"net/http"
	"net/url"
)

func CheckNeedArgs(values url.Values, args ...string) error {
	message := ""
	for _, value := range args {
		if !values.Has(value) {
			message += value + ", "
		}
	}
	if "" != message {
		return errors.New(message + "arg(s) is(are) not defined.")
	}
	return nil
}

type GetRequest struct {
	Query url.Values `json:"query"`
	Err   error      `json:"err"`
}

func DealGetRequest(w http.ResponseWriter, r *http.Request, args ...string) GetRequest {
	header.SetDefaultHeaders(w)
	request := GetRequest{
		Query: r.URL.Query(),
		Err:   nil,
	}
	if "GET" != r.Method {
		if "OPTIONS" != r.Method {
			request.Err = errors.New("Request Must be 'Get'. ")
			response.StringResponse{
				Status:  response.HttpResponseFalse,
				Message: "Server Error. " + request.Err.Error(),
				Code:    http.StatusMethodNotAllowed,
			}.Send(w, r)
		} else {
			request.Err = errors.New("OPTIONS")
			response.SendResponseOK(w, r)
		}
		return request
	}
	request.Err = CheckNeedArgs(request.Query, args...)
	if nil != request.Err {
		response.StringResponse{
			Status:  response.HttpResponseFalse,
			Message: "Server Error. " + request.Err.Error(),
			Code:    http.StatusExpectationFailed,
		}.Send(w, r)
	}
	return request
}

func DealGetError() {}
