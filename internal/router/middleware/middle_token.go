package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
)

func (m *middleware) Token(ctx core.Context) (userId int64, userName string, err errno.Error) {
	token := ctx.GetHeader("Token")
	if token == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Missing Token parameter in Header"))

		return
	}

	if !m.cache.Exists(m.adminService.CacheKeyPrefix() + token) {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Please log in first"))

		return
	}

	cacheData, cacheErr := m.cache.Get(m.adminService.CacheKeyPrefix()+token, cache.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(cacheErr)

		return
	}

	type userInfo struct {
		Id       int64  `json:"id"`       // User ID
		Username string `json:"username"` // username
	}

	var userData userInfo
	_ = json.Unmarshal([]byte(cacheData), &userData)

	userId = userData.Id
	userName = userData.Username

	return
}
