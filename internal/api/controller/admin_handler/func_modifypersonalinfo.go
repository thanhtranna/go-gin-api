package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type modifyPersonalInfoRequest struct {
	Nickname string `form:"nickname"` // Nickname
	Mobile   string `form:"mobile"`   // Mobile phone number
}

type modifyPersonalInfoResponse struct {
	Username string `json:"username"` // User account
}

// ModifyPersonalInfo modify personal information
// @Summary modify personal information
// @Description modify personal information
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param nickname formData string true "nickname"
// @Param mobile formData string true "mobile number"
// @Success 200 {object} modifyPersonalInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
func (h *handler) ModifyPersonalInfo() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyPersonalInfoRequest)
		res := new(modifyPersonalInfoResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		userId := cast.ToInt32(c.UserID())

		modifyData := new(admin_service.ModifyData)
		modifyData.Nickname = req.Nickname
		modifyData.Mobile = req.Mobile

		if err := h.adminService.ModifyPersonalInfo(c, userId, modifyData); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminModifyPersonalInfoError,
				code.Text(code.AdminModifyPersonalInfoError)).WithErr(err),
			)
			return
		}

		res.Username = c.UserName()
		c.Payload(res)
	}
}
