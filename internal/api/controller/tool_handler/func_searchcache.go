package tool_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type searchCacheRequest struct {
	RedisKey string `form:"redis_key"` // Redis Key
}

type searchCacheResponse struct {
	Val string `json:"val"` // value after query
	TTL string `json:"ttl"` // Expiration time
}

// SearchCache query cache
// @Summary query cache
// @Description query cache
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param redis_key formData string true "Redis Key"
// @Success 200 {object} searchCacheResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/cache/search [post]
func (h *handler) SearchCache() core.HandlerFunc {
	return func(c core.Context) {
		req := new(searchCacheRequest)
		res := new(searchCacheResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if b := h.cache.Exists(req.RedisKey); !b {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SearchRedisEmpty,
				code.Text(code.SearchRedisEmpty)),
			)
			return
		}

		val, err := h.cache.Get(req.RedisKey, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SearchRedisError,
				code.Text(code.SearchRedisError)).WithErr(err),
			)
			return
		}

		ttl, err := h.cache.TTL(req.RedisKey)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SearchRedisError,
				code.Text(code.SearchRedisError)).WithErr(err),
			)
			return
		}

		res.Val = val
		res.TTL = ttl.String()
		c.Payload(res)
	}
}
