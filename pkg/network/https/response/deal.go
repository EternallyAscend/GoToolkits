package response

import "net/http"

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
		if nil != err {
			return w, r, err
		}
	}
	return wc, rc, nil
}
