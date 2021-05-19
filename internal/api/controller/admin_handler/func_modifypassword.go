package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type modifyPasswordRequest struct {
	OldPassword string `form:"old_password"` // Old Password
	NewPassword string `form:"new_password"` // New Password
}

type modifyPasswordResponse struct {
	Username string `json:"username"` // Username
}

// ModifyPassword modify password
// @Summary modify password
// @Description modify password
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param old_password formData string true "old password"
// @Param new_password formData string true "new password"
// @Success 200 {object} modifyPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyPasswordRequest)
		res := new(modifyPasswordResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		userId := cast.ToInt32(c.UserID())

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Id = userId
		searchOneData.Password = password.GeneratePassword(req.OldPassword)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil || info == nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithErr(err),
			)
			return
		}

		if err := h.adminService.ModifyPassword(c, userId, req.NewPassword); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithErr(err),
			)
			return
		}

		res.Username = c.UserName()
		c.Payload(res)
	}
}
