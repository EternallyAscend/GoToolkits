package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Status  int             `json:"status"`
	Message json.RawMessage `json:"message"`
	Code    int             `json:"code"`
}

func SendJsonResponse(w http.ResponseWriter, r *http.Request, response JsonResponse) {
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

func DealJsonMarshal(w http.ResponseWriter, r *http.Request, v interface{}) []byte {
	data, err := json.Marshal(v)
	if nil != err {
		SendStringResponse(w, r, StringResponse{
			Status:  HttpResponseFalse,
			Message: "Server Error. Json: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return nil
	}
	return data
}

func WriteJsonMarshal(w http.ResponseWriter, r *http.Request, v interface{}) {
	message := DealJsonMarshal(w, r, v)
	if nil != message {
		SendJsonResponse(w, r, JsonResponse{
			Status:  HttpResponseTrue,
			Message: message,
			Code:    http.StatusOK,
		})
	}
}
