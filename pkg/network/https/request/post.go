package request

import (
	"errors"
	"github.com/EternallyAscend/GoToolkit/pkg/network/https/header"
	"github.com/EternallyAscend/GoToolkit/pkg/network/https/response"
	"net/http"
)

type PostRequest struct {
	Err error `json:"err"`
}

func DealPostRequest(w http.ResponseWriter, r *http.Request, args ...string) PostRequest {
	header.SetDefaultHeaders(w)
	request := PostRequest{
		Err: r.ParseForm(),
	}
	if "POST" != r.Method {
		if "OPTIONS" != r.Method {
			request.Err = errors.New("Request Must be 'Post'. ")
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
	request.Err = CheckNeedArgs(r.Form, args...)
	if nil != request.Err {
		response.StringResponse{
			Status:  response.HttpResponseFalse,
			Message: "Server Error. " + request.Err.Error(),
			Code:    http.StatusExpectationFailed,
		}.Send(w, r)
	}
	return request
}

func DealPostError() {}

type Dealer struct {
	Header   func(w http.ResponseWriter) `json:"header"`
	Handlers []Handler                   `json:"handlers"`
}

type Handler interface {
	Dealer(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)
}

func (dealer Dealer) Deal(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	if nil != dealer.Header {
		dealer.Header(w)
	}
	wc, rc := w, r
	var err error
	for i := range dealer.Handlers {
		if nil == dealer.Handlers[i].Dealer {
			continue
		}
		wc, rc, err = dealer.Handlers[i].Dealer(wc, rc)
		if err != nil {
			return w, r, err
		}
	}
	return wc, rc, nil
}
