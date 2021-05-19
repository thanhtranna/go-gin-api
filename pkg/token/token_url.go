package token

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// UrlSign
// path The requested path (without querystring)
func (t *token) UrlSign(path string, method string, params url.Values) (tokenString string, err error) {
	// Legal Methods
	methods := map[string]bool{
		"get":     true,
		"post":    true,
		"put":     true,
		"path":    true,
		"delete":  true,
		"head":    true,
		"options": true,
	}

	methodName := strings.ToLower(method)
	if !methods[methodName] {
		err = errors.New("method param error")
		return
	}

	// Encode() Comes in the method sorted by key
	sortParamsEncode := params.Encode()

	// Encrypted string rules path + method + sortParamsEncode + secret
	encryptStr := fmt.Sprintf("%s%s%s%s", path, methodName, sortParamsEncode, t.secret)

	// Encrypted string md5
	s := md5.New()
	s.Write([]byte(encryptStr))
	md5Str := hex.EncodeToString(s.Sum(nil))

	// Base64 encode md5Str
	tokenString = base64.StdEncoding.EncodeToString([]byte(md5Str))

	return
}
