package RBAC

import (
	"fmt"
	"net/http"
)

// Example for RBAC.

type HandlerRBAC struct {
	// Roles Allow access roles.
	Roles []string `json:"roles"`
	// Interfaces
	//RequestHandler `json:"requestHandler"` // Inherit Dealer Handler.
}

func (this HandlerRBAC) Dealer(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	fmt.Println(this.Roles)
	return w, r, nil
}

func RBAC(w http.ResponseWriter, r *http.Request) {
	//rbac := HandlerRBAC{
	//	Roles: []string{"Admin"},
	//}
	//wc, rc, err := RequestDealer{
	//	Header: nil,
	//	Handlers: []RequestHandler{
	//		rbac,
	//	},
	//}.Deal(w, r)
	//fmt.Println(wc, rc, err)
}
