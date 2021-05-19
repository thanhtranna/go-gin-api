package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	"github.com/xinliangnote/go-gin-api/pkg/time_parse"

	"github.com/pkg/errors"
)

// Generate
// path The requested path (without querystring)
func (s *signature) Generate(path string, method string, params url.Values) (authorization, date string, err error) {
	if path == "" {
		err = errors.New("path required")
		return
	}

	if method == "" {
		err = errors.New("method required")
		return
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = errors.New("method param error")
		return
	}

	// Date
	date = time_parse.CSTLayoutString()

	// Encode() Comes in the method sorted by key
	sortParamsEncode, err := url.QueryUnescape(params.Encode())
	if err != nil {
		err = errors.Errorf("url QueryUnescape error %v", err)
		return
	}

	// Encrypted string rules
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(path)
	buffer.WriteString(delimiter)
	buffer.WriteString(methodName)
	buffer.WriteString(delimiter)
	buffer.WriteString(sortParamsEncode)
	buffer.WriteString(delimiter)
	buffer.WriteString(date)

	// Perform sha256 encryption on the data and base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	authorization = fmt.Sprintf("%s %s", s.key, digest)
	return
}
