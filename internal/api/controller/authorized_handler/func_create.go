package authorized_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/authorized_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createRequest struct {
	BusinessKey       string `form:"business_key"`       // caller key
	BusinessDeveloper string `form:"business_developer"` // caller docking person
	Remark            string `form:"remark"`             // Remark
}

type createResponse struct {
	Id int32 `json:"id"` // Primary Key ID
}

// Create new caller
// @Summary add caller
// @Description add caller
// @Tags API.authorized
// @Accept multipart/form-data
// @Produce json
// @Param business_key formData string true "caller key"
// @Param business_developer formData string true "caller docking person"
// @Param remark formData string true "Remarks"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [post]
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

		createData := new(authorized_service.CreateAuthorizedData)
		createData.BusinessKey = req.BusinessKey
		createData.BusinessDeveloper = req.BusinessDeveloper
		createData.Remark = req.Remark

		id, err := h.authorizedService.Create(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
