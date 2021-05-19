package grpc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// Demonstrate the use of sign, the actual situation is subject to the signature algorithm agreed by the gRPC server

const (
	// ProxyAuthorization used by signature, both gateway and grpc
	ProxyAuthorization = "proxy-authorization"
)

type Sign func(message []byte) (auth string, err error)

func GenerateSign(secret string, message []byte) (auth string, err error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(message)

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return digest, nil
}
