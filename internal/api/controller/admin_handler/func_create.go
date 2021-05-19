package admin_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/admin_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createRequest struct {
	Username string `form:"username"` // username
	Nickname string `form:"nickname"` // Nickname
	Mobile   string `form:"mobile"`   // mobile phone number
	Password string `form:"password"` // Password
}

type createResponse struct {
	Id int32 `json:"id"` // Primary key ID
}

// Create new administrator
// @Summary add administrator
// @Description add administrator
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param nickname formData string true "nickname"
// @Param mobile formData string true "mobile number"
// @Param password formData string true "password"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		createData := new(admin_service.CreateAdminData)
		createData.Nickname = req.Nickname
		createData.Username = req.Username
		createData.Mobile = req.Mobile
		createData.Password = req.Password

		id, err := h.adminService.Create(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
