package router

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	Path     string   `json:"path"`
	Port     int      `json:"port"`
	Mask     string   `json:"mask"`
	CertPath string   `json:"certPath"`
	KeyPath  string   `json:"keyPath"`
	Routers  []Router `json:"routers"`
}

func (server Server) Listen() {
	sm := http.NewServeMux()
	for i := range server.Routers {
		server.Routers[i].attachServer(sm, server.Mask)
	}
	err := MuxListen(server.Path, server.Port, sm)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func (server Server) ListenTLS() {
	sm := http.NewServeMux()
	for i := range server.Routers {
		server.Routers[i].attachServer(sm, server.Mask)
	}
	err := MuxListenTLS(server.Path, server.Port, server.CertPath, server.KeyPath, sm)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func MuxListen(path string, port int, mux *http.ServeMux) error {
	return http.ListenAndServe(path+":"+strconv.Itoa(port), mux)
}

func MuxListenTLS(path string, port int, cert string, key string, mux *http.ServeMux) error {
	return http.ListenAndServeTLS(path+":"+strconv.Itoa(port), cert, key, mux)
}
