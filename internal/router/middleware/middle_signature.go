package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/signature"
	"github.com/xinliangnote/go-gin-api/pkg/urltable"

	"github.com/pkg/errors"
)

const ttl = time.Minute * 2 // Signature timeout 2 minutes

var whiteListPath = map[string]bool{
	"/login/web": true,
}

func (m *middleware) Signature() core.HandlerFunc {
	return func(c core.Context) {
		// Signature information
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("The Authorization parameter is missing in the header")),
			)
			return
		}

		// Time information
		date := c.GetHeader("Authorization-Date")
		if date == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Date parameter missing in header")),
			)
			return
		}

		// Obtain the key from the signature information
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Authorization format error in header")),
			)
			return
		}

		key := authorizationSplit[0]

		data, err := m.authorizedService.DetailByKey(c, key)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(err),
			)
			return
		}

		// Verify that cache is called
		if data.IsUsed == -1 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New(key + " has been banned")),
			)
			return
		}

		if len(data.Apis) < 1 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New(key + " no interface authorization")),
			)
			return
		}

		if !whiteListPath[c.Path()] {
			// Verify that c.Method() + c.Path() is authorized
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.SignatureError,
					code.Text(code.SignatureError)).WithErr(errors.New(c.Method() + c.Path() + " no interface authorization")),
				)
				return
			}
		}

		ok, err := signature.New(key, data.Secret, ttl).Verify(authorization, date, c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(err),
			)
			return
		}

		if !ok {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Authorization information in the header is wrong")),
			)
			return
		}
	}
}
