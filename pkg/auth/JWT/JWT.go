package JWT

import (
	"github.com/golang-jwt/jwt"
	"net/http"
)

// Example for Jwt.

func DecodeJWT(seg string) ([]byte, error) {
	return jwt.DecodeSegment(seg)
}

// JWT Restful https://mojotv.cn/go/golang-jwt-auth

type HandlerJsonWebToken struct {
	// Interfaces
	//RequestHandler `json:"requestHandler"` // Inherit Dealer Handler.
}

func (this HandlerJsonWebToken) Dealer(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	return w, r, nil
}
