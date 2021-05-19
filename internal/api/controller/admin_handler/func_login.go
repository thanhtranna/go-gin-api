package admin_handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/pkg/errors"
)

type loginRequest struct {
	Username string `form:"username"` // Username
	Password string `form:"password"` // Password
}

type loginResponse struct {
	Token string `json:"token"` // User token
}

// Login Administrator login
// @Summary Administrator login
// @Description Administrator login
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "password"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Username = req.Username
		searchOneData.Password = password.GeneratePassword(req.Password)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		if info == nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(errors.New("No eligible users found")),
			)
			return
		}

		token := password.GenerateLoginToken(info.Id)

		// User Info
		adminJsonInfo, _ := json.Marshal(info)

		// Save into Redis
		err = h.cache.Set(h.adminService.CacheKeyPrefix()+token, string(adminJsonInfo), time.Hour*24, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		res.Token = token
		c.Payload(res)
	}
}
