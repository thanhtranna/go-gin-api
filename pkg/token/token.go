package token

import (
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _ Token = (*token)(nil)

type Token interface {
	// i In order to avoid being implemented by other packages
	i()

	// JWT Signature method
	JwtSign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error)
	JwtParse(tokenString string) (*claims, error)

	// URL Signature method, decryption is not supported
	UrlSign(path string, method string, params url.Values) (tokenString string, err error)
}

type token struct {
	secret string
}

type claims struct {
	UserID   int64
	UserName string
	jwt.StandardClaims
}

func New(secret string) Token {
	return &token{
		secret: secret,
	}
}

func (t *token) i() {}
