package errno

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	// i in order to avoid being implemented by other packages
	i()
	// WithErr set error message
	WithErr(err error) Error
	// GetBusinessCode to get Business Code
	GetBusinessCode() int
	// GetHttpCode to get HTTP Code
	GetHttpCode() int
	// GetMsg gets Msg
	GetMsg() string
	// GetErr to get error information
	GetErr() error
	// ToString returns the error details in JSON format
	ToString() string
}

type err struct {
	HttpCode     int    // HTTP Code
	BusinessCode int    // Business Code
	Message      string // description information
	Err          error  // error message
}

func NewError(httpCode, businessCode int, msg string) Error {
	return &err{
		HttpCode:     httpCode,
		BusinessCode: businessCode,
		Message:      msg,
	}
}

func (e *err) i() {}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	return e.Message
}

func (e *err) GetErr() error {
	return e.Err
}

// ToString return error details in JSON format
func (e *err) ToString() string {
	err := &struct {
		HttpCode     int    `json:"http_code"`
		BusinessCode int    `json:"business_code"`
		Message      string `json:"message"`
	}{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessCode,
		Message:      e.Message,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
