package router

import "net/http"

type Method struct {
	Pattern string                                                  `json:"pattern"`
	Handler func(writer http.ResponseWriter, request *http.Request) `json:"handler"`
}

func MethodGenerator(pattern string, handler func(writer http.ResponseWriter, request *http.Request)) Method {
	return Method{
		Pattern: pattern,
		Handler: handler,
	}
}

type Router struct {
	Name    string   `json:"name"`
	Mask    string   `json:"mask"`
	Methods []Method `json:"methods"`
}

func (router Router) attachServer(sm *http.ServeMux, mask string) {
	for i := range router.Methods {
		sm.HandleFunc(mask+router.Mask+router.Methods[i].Pattern, router.Methods[i].Handler)
	}
}
